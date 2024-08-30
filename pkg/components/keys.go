package components

type PubKey interface {
	Bytes() []byte
	Equals(other PubKey) bool
	Type() string
}

// PrivKey interface with generics
type PrivKey[T PubKey] interface {
	Bytes() []byte
	PubKey() T
	Equals(other PrivKey[T]) bool
	Type() string
}
