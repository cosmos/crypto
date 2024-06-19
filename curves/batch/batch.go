package batch

import (
	"github.com/cosmos/crypto/curves/batch/verifier"
	"github.com/cosmos/crypto/curves/ed25519"
	"github.com/cosmos/crypto/curves/sr25519"
	"github.com/cosmos/crypto/types"
)

// CreateBatchVerifier checks if a key type implements the batch verifier interface.
// Currently only ed25519 & sr25519 supports batch verification.
func CreateBatchVerifier(pk types.PubKey) (verifier.BatchVerifier, bool) {
	switch pk.Type() {
	case ed25519.KeyType:
		return ed25519.NewBatchVerifier(), true
	case sr25519.KeyType:
		return sr25519.NewBatchVerifier(), true
	}

	// case where the key does not support batch verification
	return nil, false
}

// SupportsBatchVerifier checks if a key type implements the batch verifier
// interface.
func SupportsBatchVerifier(pk types.PubKey) bool {
	if pk == nil {
		return false
	}
	switch pk.Type() {
	case ed25519.KeyType, sr25519.KeyType:
		return true
	}

	return false
}
