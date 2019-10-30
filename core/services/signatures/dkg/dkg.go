// Package dkg administers collaboration between a set of nodes in constructing a threshold distributed key.
package dkg

import (
	"fmt"

	"github.com/smartcontractkit/chainlink/core/store/models"
	"go.dedis.ch/kyber/v3"
)

type SecretShare struct {
	Index uint
	Share kyber.Scalar
}

type peers = []models.CompressedPubKey

type SharedKey struct {
	Peers     peers
	Threshold int
	PublicKey kyber.Point
	Shares    []struct {
		Index int
		Share kyber.Scalar
	}
}

// DKGParams represents the arguments to GenerateSharedKey. See its docstring
type DKGParams struct {
	Threshold uint
	Peers     peers
	KeyIdx    uint
	SecretKey *models.PrivateKey
}

// check sanity-checks p prior to starting the DKG process
func (p DKGParams) check() error {
	if p.Threshold < 1 || p.Threshold > uint(len(p.Peers)) {
		return fmt.Errorf("threshold %d is not in {1, ..., %d}", p.Threshold, len(p.Peers))
	}
	compressed := p.SecretKey.CompressedPublicKey()
	if p.Peers[p.KeyIdx] != compressed {
		return fmt.Errorf("wrong index for self's key in participant list")
	}
	var seenPeers map[models.CompressedPubKey]bool
	for _, peer := range p.Peers {
		if seenPeers[peer] {
			return fmt.Errorf("peer %s has multiple entries on the peer list", peer)
		}
		seenPeers[peer] = true
	}
	return nil
}

// GenerateSharedKey initiates a DKG process, collaborating with the given Peers
// to construct a threshold key which can be used by any subset of Peers of size
// Threshold, and using SecretKey as its identity throughout the process, which
// must correspond to the KeyIdx'th peer on the list
func GenerateSharedKey(p DKGParams) (*SharedKey, error) {
	if err := p.check(); err != nil {
		return nil, err
	}
	peerAddresses, err := findPeers(p.Peers, p.SecretKey)
	copy(peerAddresses[:], peerAddresses) // XXX: No op
	if err != nil {
		return nil, err
	}
	return nil, nil // XXX:
}
