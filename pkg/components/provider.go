package components

// CryptoProvider aggregates the functionalities of signing, verifying, and hashing, and provides metadata.
type CryptoProvider interface {
	// GetSigner returns an instance of Signer.
	GetSigner() Signer

	// GetVerifier returns an instance of Verifier.
	GetVerifier() Verifier

	// GetHasher returns an instance of Hasher.
	GetHasher() Hasher

	// Metadata returns metadata for the crypto provider.
	Metadata() ProviderMetadata
}
