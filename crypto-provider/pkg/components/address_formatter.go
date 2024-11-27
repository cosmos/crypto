package components

// AddressFormatter returns a formatted address from the given public key bytes
type AddressFormatter interface {
	FormatAddress([]byte) (string, error)
}
