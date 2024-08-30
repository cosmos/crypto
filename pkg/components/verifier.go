package components

// Verifier represents a general interface for verifying signatures.
type Verifier interface {
	// Verify checks the digital signature against the message and a public key to determine its validity.
	Verify(signature Signature, signDoc []byte, options VerifierOpts) (bool, error)
}

type VerifierOpts = map[string]any
