package bulletprooftxmanager

import (
	"strings"

	"github.com/pkg/errors"
)

// fatal means this transaction can never be accepted even with a different nonce or higher gas price
type sendError struct {
	fatal bool
	err   error
}

func (f *sendError) Error() string {
	return f.err.Error()
}

func (f *sendError) StrPtr() *string {
	e := f.err.Error()
	return &e
}

func (s *sendError) Fatal() bool {
	return s != nil && s.fatal
}

// Geth/parity returns this error if a transaction with this nonce already
// exists either on-chain or in the mempool.
//
// There are two scenarios in which this can happen:
// 1. The private key has been used to send at least one transaction from another wallet
// 2. The chainlink node crashed before being able to save the broadcastAt timestamp, indicating
//    that we are trying to send the exact same transaction twice (but it was already mined into a block).
//
// We can know which it is, because if we crashed there will be an unfinishedEthTransaction in the database.
// TODO: Probably needs a unit test
func (s *sendError) isNonceAlreadyUsedError() bool {
	// TODO: Add parity error
	return s != nil && s.err != nil && (s.err.Error() == "nonce too low" || s.err.Error() == "replacement transaction underpriced")
}

// Geth/parity returns this error if the transaction is already in the node's mempool
func (s *sendError) isTransactionAlreadyInMempool() bool {
	// TODO: Needs parity errors here
	return s.err != nil && strings.HasPrefix(s.Error(), "known transaction:")
}

// TODO: Write doc
func (s *sendError) isTerminallyUnderpriced() bool {
	// TODO: geth/parity errors
	return s.err != nil && (s.Error() == "transaction underpriced")
}

func NewFatalSendError(s string) *sendError {
	return &sendError{err: errors.New(s), fatal: true}
}

func FatalSendError(e error) *sendError {
	if e == nil {
		return nil
	}
	return &sendError{err: e, fatal: true}
}

func SendError(e error) *sendError {
	if e == nil {
		return nil
	}
	fatal := isFatalSendError(e)
	return &sendError{err: e, fatal: fatal}
}

// Geth/parity returns these errors if the transaction failed in such a way that:
// 1. It can NEVER be included into a block
// 2. Resending the transaction even with higher gas price will never change that outcome
// TODO: This probably should have unit tests
func isFatalSendError(err error) bool {
	if err == nil {
		return false
	}
	switch err.Error() {
	// Geth errors
	// See: https://github.com/ethereum/go-ethereum/blob/b9df7ecdc3d3685180ceb29665bab59e9f614da5/core/tx_pool.go#L516
	case "exceeds block gas limit", "invalid sender", "negative value", "oversized data", "gas uint64 overflow", "intrinsic gas too low", "nonce too high":
		return true
	// TODO: Add parity here, and can we use error codes?
	// See: https://github.com/openethereum/openethereum/blob/master/rpc/src/v1/helpers/errors.rs#L420
	default:
		return false
	}
}
