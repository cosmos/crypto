package file

import "github.com/cosmos/crypto-provider/pkg/components"

type (
	PubKey  = components.PubKey
	PrivKey = components.PrivKey[components.PubKey]
)

// dummyPubKey implements the PubKey interface
type dummyPubKey struct {
	key string
}

func (d *dummyPubKey) Bytes() []byte {
	return []byte(d.key)
}

func (d *dummyPubKey) Equals(other components.PubKey) bool {
	if op, ok := other.(*dummyPubKey); ok {
		return d.key == op.key
	}
	return false
}

func (d *dummyPubKey) Type() string {
	return "dummyPubKey"
}

// NewPubKeyFromString creates a new PubKey instance from a string
func NewPubKeyFromString(key string) components.PubKey {
	return &dummyPubKey{key: key}
}

// Ensure dummyPrivKey implements the PrivKey interface
var _ PrivKey = (*dummyPrivKey)(nil)

// Ensure dummyPubKey implements the PubKey interface
var _ components.PubKey = (*dummyPubKey)(nil)

type dummyPrivKey struct {
	key string
}

func (d *dummyPrivKey) Bytes() []byte {
	return []byte(d.key)
}

func (d *dummyPrivKey) PubKey() components.PubKey {
	// Create public key from private key
	return NewPubKeyFromString(d.key)
}

func (d *dummyPrivKey) Equals(other components.PrivKey[components.PubKey]) bool {
	if op, ok := other.(*dummyPrivKey); ok {
		return d.key == op.key
	}
	return false
}

func (d *dummyPrivKey) Type() string {
	return "dummyPrivKey"
}
