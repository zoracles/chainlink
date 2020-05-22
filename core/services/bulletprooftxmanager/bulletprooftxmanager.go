package bulletprooftxmanager

import (
	"bytes"
	"context"
	"math/big"
	"time"

	"github.com/smartcontractkit/chainlink/core/eth"
	"github.com/smartcontractkit/chainlink/core/logger"
	strpkg "github.com/smartcontractkit/chainlink/core/store"
	"github.com/smartcontractkit/chainlink/core/store/models"
	"github.com/smartcontractkit/chainlink/core/store/orm"
	"github.com/smartcontractkit/chainlink/core/utils"

	gethAccounts "github.com/ethereum/go-ethereum/accounts"
	gethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

const (
	// maxEthNodeRequestTime is the worst case time we will wait for a response
	// from the eth node before we consider it to be an error
	maxEthNodeRequestTime = 2 * time.Minute
)

// TODO: Write this doc
// NOTE: it can modify the EthTransaction and the EthTransactionAttempt in
// memory but will not save them
// Returning error here indicates that it may succeed on retry
func send(store *strpkg.Store, etx *models.EthTransaction, attempt *models.EthTransactionAttempt, initialGasPrice *big.Int) *sendError {
	if etx == nil || attempt == nil {
		return NewFatalSendError("etx and etxAttempt must be non-nil")
	}
	if etx.Nonce == nil {
		return NewFatalSendError("cannot send transaction without nonce")
	}
	account, err := store.KeyStore.GetAccountByAddress(etx.FromAddress)
	if err != nil {
		return FatalSendError(errors.Wrapf(err, "Error getting account %s for transaction %v", etx.FromAddress.String(), etx.ID))
	}

	transaction := gethTypes.NewTransaction(uint64(*etx.Nonce), etx.ToAddress, etx.Value.ToInt(), etx.GasLimit, initialGasPrice, etx.EncodedPayload)
	signedTx, signedTxBytes, err := signTx(store.KeyStore, account, transaction, store.Config.ChainID())
	if err != nil {
		return FatalSendError(errors.Wrapf(err, "Error using account %s to sign transaction %v", etx.FromAddress.String(), etx.ID))
	}

	attempt.SignedRawTx = signedTxBytes
	attempt.EthTransactionID = etx.ID
	attempt.GasPrice = *utils.NewBig(initialGasPrice)
	attempt.Hash = transaction.Hash()

	sendErr := sendTransaction(store.GethClientWrapper, signedTx)
	broadcastAt := time.Now()

	if sendErr.Fatal() {
		return sendErr
	}

	etx.BroadcastAt = &broadcastAt

	if sendErr == nil {
		return nil
	}

	// Bump gas if necessary
	if sendErr.isTerminallyUnderpriced() {
		logger.Errorf("transaction %v was underpriced at %v wei. You should increase your configured ETH_GAS_PRICE_DEFAULT (currently set to %v wei)", etx.ID, initialGasPrice, store.Config.EthGasPriceDefault())
		newGasPrice := BumpGas(store.Config, initialGasPrice)
		logger.Infof("retrying transaction %v with new gas price of %v wei", etx.ID, newGasPrice.Int64())
		return send(store, etx, attempt, newGasPrice)
	} else if sendErr.isTransactionAlreadyInMempool() {
		logger.Debugf("transaction %v already in mempool", etx.ID)
		return nil
	}
	return sendErr
}

func signTx(keyStore strpkg.KeyStoreInterface, account gethAccounts.Account, tx *gethTypes.Transaction, chainID *big.Int) (*gethTypes.Transaction, []byte, error) {
	signedTx, err := keyStore.SignTx(account, tx, chainID)
	if err != nil {
		return nil, nil, err
	}
	rlp := new(bytes.Buffer)
	if err := signedTx.EncodeRLP(rlp); err != nil {
		return nil, nil, err
	}
	return signedTx, rlp.Bytes(), nil

}

func sendTransaction(gethClientWrapper strpkg.GethClientWrapper, signedTransaction *gethTypes.Transaction) *sendError {
	err := gethClientWrapper.GethClient(func(gethClient eth.GethClient) error {
		ctx, cancel := context.WithTimeout(context.Background(), maxEthNodeRequestTime)
		defer cancel()
		return gethClient.SendTransaction(ctx, signedTransaction)
	})

	return SendError(err)
}

// TODO: This is copied (with changes - need to integrate) from tx_manager which is suboptimal. Consider copying unit tests also.
// bumpGas returns a new gas price increased by the larger of:
// - A configured percentage bump (ETH_GAS_BUMP_PERCENT)
// - A configured fixed amount of Wei (ETH_GAS_PRICE_WEI)
// TODO: unit test
func BumpGas(config orm.ConfigReader, originalGasPrice *big.Int) *big.Int {
	// Similar logic is used in geth
	// See: https://github.com/ethereum/go-ethereum/blob/8d7aa9078f8a94c2c10b1d11e04242df0ea91e5b/core/tx_list.go#L255
	// And: https://github.com/ethereum/go-ethereum/blob/8d7aa9078f8a94c2c10b1d11e04242df0ea91e5b/core/tx_pool.go#L171
	percentageMultiplier := big.NewInt(100 + int64(config.EthGasBumpPercent()))
	minimumGasBumpByPercentage := new(big.Int).Div(
		new(big.Int).Mul(
			originalGasPrice,
			percentageMultiplier,
		),
		big.NewInt(100),
	)
	minimumGasBumpByIncrement := new(big.Int).Add(originalGasPrice, config.EthGasBumpWei())
	currentDefaultGasPrice := config.EthGasPriceDefault()
	prices := []*big.Int{minimumGasBumpByPercentage, minimumGasBumpByIncrement, currentDefaultGasPrice}
	max := utils.BigIntSlice(prices).Max()
	if max.Cmp(config.EthMaxGasPriceWei()) > 0 {
		logger.Errorf("bumped gas price of %v would exceed configured ETH_MAX_GAS_PRICE_WEI, capping at %v wei", max, config.EthMaxGasPriceWei())
		return config.EthMaxGasPriceWei()
	}
	return max
}
