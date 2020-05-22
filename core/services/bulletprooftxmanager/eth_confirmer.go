package bulletprooftxmanager

import (
	"context"

	"github.com/smartcontractkit/chainlink/core/eth"
	"github.com/smartcontractkit/chainlink/core/logger"
	"github.com/smartcontractkit/chainlink/core/store"
	"github.com/smartcontractkit/chainlink/core/store/models"
	"github.com/smartcontractkit/chainlink/core/utils"

	gethCommon "github.com/ethereum/go-ethereum/common"
	gethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/jinzhu/gorm"
)

type EthConfirmer interface {
	store.HeadTrackable
}

type ethConfirmer struct {
	store             *store.Store
	gethClientWrapper store.GethClientWrapper
}

func (ec *ethConfirmer) Connect(*models.Head) error {
	// TODO
	return nil
}

func (ec *ethConfirmer) Disconnect() {
	// TODO
	return
}

// TODO: Concurrent safety - advisory lock like eth_broadcaster?
func (ec *ethConfirmer) OnNewHead(head *models.Head) {
	if head == nil {
		return
	}
	if err := ec.processHead(*head); err != nil {
		logger.Error(err)
	}
}

func (ec *ethConfirmer) processHead(head models.Head) error {
	// TODO
	// Step 1: Set broadcast at block height for all attempts
	ec.setBroadcastBeforeBlockNum(head.Number)
	// Step 2: Check for receipts
	if err := ec.checkForReceipts(); err != nil {
		return err
	}

	// Step 3: See if any have exceeded the gas bumping block threshold and bump them
	for {
		// TODO: some kind of timing loop
		etxs, err := ec.findEthTransactionsRequiringNewAttempt(head.Number, int64(ec.store.Config.EthGasBumpThreshold()))
		if err != nil {
			return err
		}
		for _, etx := range etxs {
			// TODO: Test behaviour in case of crash
			// TODO: Test to make sure it actually bumps the most expensive one first
			// FIXME: Probably dont want `send` to modify the etx here
			attempt, err := ec.newAttemptWithGasBump(&etx, etx.EthTransactionAttempts[0].GasPrice)
			if err == nil || err.isTransactionAlreadyInMempool() {
				if err := ec.store.GetRawDB().Create(&attempt).Error; err != nil {
					return err
				}
				continue
			} else if err.Fatal() {
				// Should NEVER be fatal this is an invariant violation. The
				// EthBroadcaster can never create an EthTransactionAttempt that will
				// fatally error.
				logger.Error(err)
				continue
			} else if err.isNonceAlreadyUsedError() {
				// Nonce too low indicated that it was confirmed already. Success!
				receipt, err := ec.fetchReceipt(attempt.Hash)
				if err != nil {
					return err
				}
				if receipt != nil {
					if err := ec.saveReceipt(&attempt, *receipt); err != nil {
						return err
					}
					continue
				}
			}
			// Any other type of error is considered temporary e.g. network connection errors

			return err
		}
	}
}

func (ec *ethConfirmer) newAttemptWithGasBump(etx *models.EthTransaction, previousGasPrice utils.Big) (models.EthTransactionAttempt, *sendError) {
	// send
	attempt := models.EthTransactionAttempt{}
	bumpedGasPrice := BumpGas(ec.store.Config, previousGasPrice.ToInt())
	err := send(ec.store, etx, &attempt, bumpedGasPrice)
	return attempt, err
}

func (ec *ethConfirmer) checkForReceipts() error {
	// TODO: Check them all for receipts
	// Can we use goroutines to speed this up?
	unconfirmedEtxs, err := ec.findUnconfirmedEthTransactions()
	if err != nil {
		return err
	}
	for _, etx := range unconfirmedEtxs {
		for _, attempt := range etx.EthTransactionAttempts {
			receipt, err := ec.fetchReceipt(attempt.Hash)
			if err != nil {
				return err
			}
			if receipt != nil {
				if err := ec.saveReceipt(&attempt, *receipt); err != nil {
					return err
				}
				break
			}
		}
	}
	return nil
}

func (ec *ethConfirmer) fetchReceipt(hash gethCommon.Hash) (*gethTypes.Receipt, error) {
	var receipt *gethTypes.Receipt
	err := ec.gethClientWrapper.GethClient(func(gethClient eth.GethClient) error {
		ctx, cancel := context.WithTimeout(context.Background(), maxEthNodeRequestTime)
		defer cancel()
		var err error
		receipt, err = gethClient.TransactionReceipt(ctx, hash)
		return err
	})
	return receipt, err
}

func (ec *ethConfirmer) saveReceipt(attempt *models.EthTransactionAttempt, receipt gethTypes.Receipt) error {
	return ec.store.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&attempt).Error; err != nil {
			return err
		}
		if err := tx.Create(&models.EthReceipt{
			EthTransactionAttemptID: attempt.EthTransactionID,
			Receipt:                 receipt,
			TxHash:                  receipt.TxHash,
			BlockHash:               receipt.BlockHash,
			BlockNumber:             receipt.BlockNumber.Int64(),
			TransactionIndex:        receipt.TransactionIndex,
		}).Error; err != nil {
			return err
		}
		return tx.Exec(`UPDATE eth_transactions SET attempt_state = 'confirmed' WHERE id = ?`, attempt.EthTransactionID).Error
	})
}

func (ec *ethConfirmer) setBroadcastBeforeBlockNum(blockNum int64) error {
	return ec.store.GetRawDB().Exec(
		`UPDATE eth_transaction_attempts SET broadcast_before_block_num = ? WHERE broadcast_before_block_num IS NULL`,
		blockNum-1,
	).Error
}

func (ec *ethConfirmer) findUnconfirmedEthTransactions() ([]models.EthTransaction, error) {
	var etxs []models.EthTransaction
	err := ec.store.GetRawDB().
		Preload("EthTransactionAttempts", func(db *gorm.DB) *gorm.DB {
			return db.Order("eth_transaction_attempts.gas_price DESC")
		}).
		Order("nonce ASC").
		Find(&etxs, "eth_transactions_attempt_state = 'unconfirmed'").Error
	return etxs, err
}

// TODO: Indexes
func (ec *ethConfirmer) findEthTransactionsRequiringNewAttempt(blockNum int64, gasBumpThreshold int64) ([]models.EthTransaction, error) {
	var etxs []models.EthTransaction
	err := ec.store.GetRawDB().
		Preload("EthTransactionAttempts", func(db *gorm.DB) *gorm.DB {
			return db.Order("eth_transaction_attempts.gas_price DESC")
		}).
		Joins("JOIN eth_transaction_attempts ON eth_transactions.id = eth_transaction_attempts.eth_transaction_id").
		Order("nonce ASC").
		Where("eth_transactions_attempt_state = 'unconfirmed' AND confirmed_in_block_num IS NULL AND broadcast_before_block_num < ?", blockNum+gasBumpThreshold).
		Find(&etxs).Error

	return etxs, err
}
