package signer

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/cosmos/crypto-provider/pkg/components"
	"os"
	"path/filepath"
)

type FileSigner struct {
	privKeyPath string
}

func NewFileSigner(privKeyPath string) *FileSigner {
	return &FileSigner{privKeyPath: privKeyPath}
}

func (fs FileSigner) Sign(signDoc []byte, options components.SignerOpts) (components.Signature, error) {
	currentDir, _ := os.Getwd()
	pemData, err := os.ReadFile(filepath.Join(currentDir, fs.privKeyPath))
	if err != nil {
		return nil, fmt.Errorf("failed to read private key file: %v", err)
	}

	// Parse the JSON file
	var keyData struct {
		PrivKey struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"priv_key"`
	}

	err = json.Unmarshal(pemData, &keyData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON key file: %v", err)
	}

	// Decode the base64 private key
	privKeyBytes, err := base64.StdEncoding.DecodeString(keyData.PrivKey.Value)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64 private key: %v", err)
	}

	// Ensure the key is of the correct type and length
	if keyData.PrivKey.Type != "Ed25519" || len(privKeyBytes) != ed25519.PrivateKeySize {
		return nil, fmt.Errorf("invalid private key: expected Ed25519 key of length %d, got %s key of length %d", ed25519.PrivateKeySize, keyData.PrivKey.Type, len(privKeyBytes))
	}

	privKey := ed25519.PrivateKey(privKeyBytes)

	// Sign the document
	signature := ed25519.Sign(privKey, signDoc)

	return &FileSignature{data: signature}, nil
}

// FileSignature implements the types.Signature interface
type FileSignature struct {
	data []byte
}

func (fs *FileSignature) Bytes() []byte {
	return fs.data
}

func (fs *FileSignature) Equals(other components.Signature) bool {
	return string(fs.data) == string(other.Bytes())
}
