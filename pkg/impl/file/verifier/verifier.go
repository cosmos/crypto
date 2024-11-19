package verifier

import (
	"crypto/ed25519"
	"encoding/base64"
	"fmt"
	"github.com/cosmos/crypto-provider/pkg/components"
)

type FileSigVerifier struct {
	pubKey ed25519.PublicKey
}

func NewFileSigVerifier(pubKey string) (*FileSigVerifier, error) {
	pubKeyBytes, err := base64.StdEncoding.DecodeString(pubKey)
	if err != nil {
		return nil, err
	}
	if len(pubKeyBytes) != ed25519.PublicKeySize {
		return nil, fmt.Errorf("invalid public key size: expected %d, got %d", ed25519.PublicKeySize, len(pubKeyBytes))
	}
	return &FileSigVerifier{
		pubKey: ed25519.PublicKey(pubKeyBytes),
	}, nil
}

func (v *FileSigVerifier) Verify(signature components.Signature, signDoc []byte, options components.VerifierOpts) (bool, error) {
	sigBytes := signature.Bytes()
	return ed25519.Verify(v.pubKey, signDoc, sigBytes), nil
}
