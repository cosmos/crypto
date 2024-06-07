package batch

import (
	"cosmos-crypto/curves/ed25519"
	"cosmos-crypto/curves/sr25519"
	crypto "cosmos-crypto/types"
)

type BatchVerifier interface {
	// Add appends an entry into the BatchVerifier.
	Add(key crypto.PubKey, message, signature []byte) error
	// Verify verifies all the entries in the BatchVerifier, and returns
	// if every signature in the batch is valid, and a vector of bools
	// indicating the verification status of each signature (in the order
	// that signatures were added to the batch).
	Verify() (bool, []bool)
}

// CreateBatchVerifier checks if a key type implements the batch verifier interface.
// Currently only ed25519 & sr25519 supports batch verification.
func CreateBatchVerifier(pk crypto.PubKey) (BatchVerifier, bool) {
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
func SupportsBatchVerifier(pk crypto.PubKey) bool {
	if pk == nil {
		return false
	}
	switch pk.Type() {
	case ed25519.KeyType, sr25519.KeyType:
		return true
	}

	return false
}
