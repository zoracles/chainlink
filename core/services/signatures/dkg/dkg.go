// Package dkg administers collaboration between a set of nodes in constructing a threshold distributed key.
package dkg

import (
	"fmt"
	"regexp"

	"go.dedis.ch/kyber/v3"
)

type SharedKey struct {
	Peers     []string
	Threshold int
	PublicKey kyber.Point
	Shares    []struct {
		Index int
		Share kyber.Scalar
	}
}

var address = regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
var zeroKey [32]byte

func GenerateSharedKey(peers []string, threshold int) (*SharedKey, error) {
	for _, peer := range peers {
		if !address.Match([]byte(peer)) {
			return nil, fmt.Errorf("peer %s is not an ethereum address", peer)
		}
	}
	if threshold < 1 || threshold > len(peers) {
		return nil, fmt.Errorf("threshold %d is not in {1, ..., %d}", threshold, len(peers))
	}
	return nil, nil // XXX:
}
