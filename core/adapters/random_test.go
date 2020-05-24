package adapters_test

import (
	"math/big"
	"testing"

	"github.com/smartcontractkit/chainlink/core/adapters"
	"github.com/smartcontractkit/chainlink/core/internal/cltest"
	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/solidity_verifier_wrapper"
	"github.com/smartcontractkit/chainlink/core/services/vrf"
	"github.com/smartcontractkit/chainlink/core/store/models"
	"github.com/smartcontractkit/chainlink/core/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// NB: For changes to the VRF solidity code to be reflected here, "go generate"
// must be run in core/services/vrf.
func vrfVerifier(t *testing.T) *solidity_verifier_wrapper.VRFTestHelper {
	ethereumKey, err := crypto.GenerateKey()
	require.NoError(t, err)
	auth := bind.NewKeyedTransactor(ethereumKey)
	genesisData := core.GenesisAlloc{auth.From: {Balance: big.NewInt(1000000000)}}
	gasLimit := eth.DefaultConfig.Miner.GasCeil
	backend := backends.NewSimulatedBackend(genesisData, gasLimit)
	_, _, verifier, err := solidity_verifier_wrapper.DeployVRFTestHelper(auth, backend)
	require.NoError(t, err)
	backend.Commit()
	return verifier
}

func TestRandom_Perform(t *testing.T) {
	store, cleanup := cltest.NewStore(t)
	defer cleanup()
	publicKey := cltest.StoredVRFKey(t, store)
	adapter := adapters.Random{PublicKey: publicKey.String()}
	hash := utils.MustHash("a random string")
	blockNum := 10
	jsonInput, err := models.JSON{}.MultiAdd(models.KV{
		"seed":      "0x10",
		"keyHash":   publicKey.MustHash().Hex(),
		"blockHash": hash.Hex(),
		"blockNum":  blockNum,
	})
	require.NoError(t, err) // Can't fail
	input := models.NewRunInput(&models.ID{}, jsonInput, models.RunStatusUnstarted)
	result := adapter.Perform(*input, store)
	require.NoError(t, result.Error(), "while running random adapter")
	proof := hexutil.MustDecode(result.Result().String())
	// Check respons is a valid vrf.MarshaledOnChainResponse
	length := big.NewInt(0).SetBytes(proof[:utils.EVMWordByteLen]).Uint64()
	require.Equal(t, length, uint64(len(proof)-utils.EVMWordByteLen))
	rawOnChainResponse := proof[utils.EVMWordByteLen:]
	var onChainResponse vrf.MarshaledOnChainResponse
	require.Equal(t, copy(onChainResponse[:], rawOnChainResponse),
		vrf.OnChainResponseLength, "wrong response length")
	response, err := vrf.UnmarshalProofResponse(onChainResponse)
	require.NoError(t, err, "random adapter produced bad proof response")
	actualProof, err := response.ActualProof(hash)
	require.NoError(t, err)
	require.NoError(t, err, "could not verify proof from random adapter")
	expected := common.HexToHash(
		"0x71a7c50918feaa753485ae039cb84ddd70c5c85f66b236138dea453a23d0f27e")
	assert.Equal(t, expected, common.BigToHash(actualProof.Output),
		"unexpected VRF output; perhas vrfkey.json or the output hashing function "+
			"in RandomValueFromVRFProof has changed?")
	jsonInput, err = jsonInput.Add("keyHash", common.Hash{})
	require.NoError(t, err)
	input = models.NewRunInput(&models.ID{}, jsonInput, models.RunStatusUnstarted)
	result = adapter.Perform(*input, store)
	require.Error(t, result.Error(), "must reject if keyHash doesn't match")
}
