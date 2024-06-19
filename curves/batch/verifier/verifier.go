package verifier

import "github.com/cosmos/crypto/types"

type BatchVerifier interface {
	// Add appends an entry into the BatchVerifier.
	Add(key types.PubKey, message, signature []byte) error
	// Verify verifies all the entries in the BatchVerifier, and returns
	// if every signature in the batch is valid, and a vector of bools
	// indicating the verification status of each signature (in the order
	// that signatures were added to the batch).
	Verify() (bool, []bool)
}
