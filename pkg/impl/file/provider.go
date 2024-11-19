package file

import (
	"github.com/cosmos/crypto-provider/pkg/components"
	"github.com/cosmos/crypto-provider/pkg/impl/file/hasher"
	"github.com/cosmos/crypto-provider/pkg/impl/file/signer"
	"github.com/cosmos/crypto-provider/pkg/impl/file/verifier"
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

func (fp *FileProvider) GetPubKey() components.PubKey {
	return NewPubKeyFromString(fp.metadata.PublicKey)
}

func (fp *FileProvider) InitializeKeys() error {
	return nil
}
