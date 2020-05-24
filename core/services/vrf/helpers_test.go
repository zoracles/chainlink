package vrf

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func GenerateProofWithNonce(secretKey, seed, nonce *big.Int) (*Proof, error) {
	return generateProofWithNonce(secretKey, seed, nonce)
}

func GenerateProofResponseWithNonce(secretKey, preSeed *big.Int,
	blockHash common.Hash, blockNum uint64, nonce *big.Int) (
	*MarshaledOnChainResponse, error) {
	seed := FinalSeed(preSeed, blockHash)
	proof, err := generateProofWithNonce(secretKey, seed, nonce)
	if err != nil {
		return nil, err
	}
	rv, err := (&ProofResponse{
		P: proof, PreSeed: preSeed, BlockNum: blockNum}).MarshalForVRFCoordinator()
	if err != nil {
		return nil, err
	}
	return rv, nil
}
