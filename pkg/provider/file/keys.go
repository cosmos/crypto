package file

import "crypto-provider/pkg/components"

type (
	PubKey  = components.PubKey
	PrivKey = components.PrivKey[components.PubKey]
)

// Ensure dummyPrivKey implements the PrivKey interface
var _ PrivKey = (*dummyPrivKey)(nil)

type dummyPrivKey struct {
	key string
}

func (d *dummyPrivKey) Bytes() []byte {
	return []byte(d.key)
}

func (d *dummyPrivKey) PubKey() components.PubKey {
	// Implement this based on your PubKey type
	return nil
}

func (d *dummyPrivKey) Equals(other components.PrivKey[components.PubKey]) bool {
	// Implement proper comparison
	return false
}

func (d *dummyPrivKey) Type() string {
	return "dummyPrivKey"
}
