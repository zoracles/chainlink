package bulletprooftxmanager_test

import (
	"math/big"
	"testing"
	"time"

	"github.com/smartcontractkit/chainlink/core/assets"
	"github.com/smartcontractkit/chainlink/core/internal/cltest"
	"github.com/smartcontractkit/chainlink/core/services/bulletprooftxmanager"
	"github.com/smartcontractkit/chainlink/core/store/models"
	"github.com/smartcontractkit/chainlink/core/utils"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	gethCommon "github.com/ethereum/go-ethereum/common"
)

func TestEthConfirmer_SetBroadcastBeforeBlockNum(t *testing.T) {
	store, cleanup := cltest.NewStore(t)
	defer cleanup()
	// Use the real KeyStore loaded from database fixtures
	store.KeyStore.Unlock(cltest.Password)

	config, cleanup := cltest.NewConfig(t)
	ec := bulletprooftxmanager.NewEthConfirmer(store, config)

	keys, err := store.Keys()
	require.NoError(t, err)
	key := keys[0]
	defaultFromAddress := key.Address.Address()
	toAddress := gethCommon.HexToAddress("0x6C03DDA95a2AEd917EeCc6eddD4b9D16E6380411")
	timeNow := time.Now()
	encodedPayload := []byte{1, 2, 3}
	value := assets.NewEthValue(142)
	gasLimit := uint64(242)
	gasPrice := utils.NewBig(big.NewInt(1000000000))
	nonce := int64(0)
	signedRawTx := hexutil.MustDecode("0xf867808504a817c80081f2946c03dda95a2aed917eecc6eddd4b9d16e6380411818e832a2a0029a0dd5cf86fe8e6c6c863c5cc4feb2cbfa5a87b289d8f74b8d82a599931629970faa01e65293571cd92fb96398dfd22362e76cacb527ff9472c5aa14439ae3381e9d2")
	hash := gethCommon.HexToHash("0xb025c9270b7d6df1d92730913e956146cb2db018a3611fd53c2d90abf7ef8a04")

	t.Run("saves block num to unconfirmed eth_transaction_attempts without one", func(t *testing.T) {
		headNum := int64(9000)

		// TODO: Move this to a fixture because it is getting out of hand
		etx := models.EthTransaction{
			FromAddress:    defaultFromAddress,
			ToAddress:      toAddress,
			EncodedPayload: encodedPayload,
			Value:          value,
			GasLimit:       gasLimit,
			Nonce:          &nonce,
			BroadcastAt:    &timeNow,
		}
		require.NoError(t, store.GetRawDB().Save(&etx).Error)
		attempt := models.EthTransactionAttempt{
			EthTransactionID: etx.ID,
			GasPrice:         *gasPrice,
			SignedRawTx:      signedRawTx,
			Hash:             hash,
		}
		require.NoError(t, store.GetRawDB().Save(&attempt).Error)

		// Do the thing
		require.NoError(t, ec.SetBroadcastBeforeBlockNum(headNum))

		etx, err = store.FindEthTransactionWithAttempts(etx.ID)
		require.NoError(t, err)
		require.Len(t, etx.EthTransactionAttempts, 1)
		attempt = etx.EthTransactionAttempts[0]

		assert.Equal(t, int64(8999), *attempt.BroadcastBeforeBlockNum)
	})
}
