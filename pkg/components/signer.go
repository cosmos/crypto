package components

// Signer represents a general interface for signing messages.
type Signer interface {
	// Sign takes a signDoc as input and returns the digital signature.
	Sign(signDoc []byte, options SignerOpts) (Signature, error)
}

type SignerOpts = map[string]any

// Signature represents a general interface for a digital signature.
type Signature interface {
	// Bytes returns the byte representation of the signature.
	Bytes() []byte

	// Equals checks if two signatures are identical.
	Equals(other Signature) bool
}
