package file

import (
	"crypto-provider/pkg/components"
	"crypto-provider/pkg/provider/file/hasher"
	"crypto-provider/pkg/provider/file/signer"
	"crypto-provider/pkg/provider/file/verifier"
)

const (
	ProviderTypeFile = "file"
	Version          = "v1.0.0"
)

type FileProvider struct {
	filePath string
	metadata components.ProviderMetadata
}

var _ components.CryptoProvider = &FileProvider{}

// GetSigner returns an instance of Signer.
func (fp *FileProvider) GetSigner() components.Signer {
	return signer.NewFileSigner(fp.filePath)
}

// GetVerifier returns an instance of Verifier.
func (fp *FileProvider) GetVerifier() components.Verifier {
	v, _ := verifier.NewFileSigVerifier(fp.metadata.PublicKey)
	return v
}

// GetHasher returns an instance of Hasher.
func (fp *FileProvider) GetHasher() components.Hasher {
	return hasher.FileHash{}
}

// Metadata returns metadata for the crypto provider.
func (fp *FileProvider) Metadata() components.ProviderMetadata {
	return fp.metadata
}
