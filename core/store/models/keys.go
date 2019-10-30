package models

// XXX: Make the private key a private field, and expose all functionality
// needed for the DKG/threshold crypto on public methods. Follow the geth policy
// of least privilege.

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type PrivateKey ecdsa.PrivateKey

// String elides the secret information from k
//
// This reduces the risk of accidentally logging it.
func (k PrivateKey) String() string {
	CheckKey(k.PublicKey)
	return fmt.Sprintf(
		"PrivateKey{PublicKey:0x%066x, Address=0x%040x, SecretKey=<elided>}",
		Compressed(k.PublicKey), crypto.PubkeyToAddress(k.PublicKey))
}

var bigZero = big.NewInt(0)

// CheckKey panics if k is not a valid secp256k1 key
func CheckKey(k ecdsa.PublicKey) {
	if k.X.Cmp(bigZero) == 0 && k.Y.Cmp(bigZero) == 0 {
		panic(fmt.Errorf("public key should not be 0!"))
	}
	if k.Curve.Params().Name != "secp256k1" {
		panic(fmt.Errorf("key must be secp256k1 point"))
	}
	if !k.Curve.IsOnCurve(k.X, k.Y) {
		panic(fmt.Errorf("key is not on secp256k1 curve!"))
	}
}

var bigOne = big.NewInt(1)

// CompressedPubKey is a binary representation of a secp256k1 key.
//
// The first byte of CompressedPubKey corresponding to curve point (x, y)
// contains 2 if y is even, 3 if odd. The remaining 32 bytes contain x's
// big-endian byte representation.
//
// See https://www.secg.org/sec1-v2.pdf, section 2.3.3, p. 10, steps 2.2-3
type CompressedPubKey [33]byte

// Compressed is hex of k in binary compressed format
func Compressed(k ecdsa.PublicKey) CompressedPubKey {
	CheckKey(k)
	yLowestBit := new(big.Int).And(k.Y, bigOne).Int64()
	if yLowestBit != 0 && yLowestBit != 1 {
		panic("failed to compute lowest bit of y")
	}
	var rv CompressedPubKey
	rv[0] = byte(yLowestBit)
	xBinary := k.X.Bytes()
	if len(xBinary) > 32 {
		panic("x too large")
	}
	copy(rv[1:], append(make([]byte, 32-len(xBinary)), xBinary...))
	return rv
}

var compressedPubKeyHex = regexp.MustCompile("^(0[xX])[0-9a-fA-F]{66}")

// FromHex stores in c the compressed public key represented by the hex in s.
// s must be a 66-nybble hex string, optionally 0x-prefixed.
func (c CompressedPubKey) FromHex(s string) error {
	if !compressedPubKeyHex.Match([]byte(s)) {
		return fmt.Errorf("Cannot parse %s as hex compressed public key", s)
	}
	copy(c[:], common.Hex2Bytes(s))
	return nil
}

// Compressed returns p's public key in compressed binary format
func (p PrivateKey) CompressedPublicKey() CompressedPubKey {
	return Compressed(p.PublicKey)
}

func libP2PId(k ecdsa.PublicKey) {}
