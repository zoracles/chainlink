// Package dkg administers collaboration between a set of nodes in constructing a threshold distributed key.
package dkg

type SharedKey struct {
	Peers string []
	Threshold int
	
}

func GenerateSharedKey(peers []string, threshold int) (SharedKey, error) {

}
