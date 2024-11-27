package components

// Hasher represents a general interface for hashing data.
type Hasher interface {
	// Hash takes an input byte array and returns the hashed output as a byte array.
	Hash(input []byte, options HasherOpts) (output []byte, err error)
}

type HasherOpts = map[string]any
