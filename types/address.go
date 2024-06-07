package types

import (
	"cosmos-crypto/internal/libs/bytes"
)

const (
	// AddressSize is the size of a pubkey address.
	AddressSize = hash.TruncatedSize
)

// Address An address is a []byte, but hex-encoded even in JSON.
// []byte leaves us the option to change the address length.
// Use an alias so Unmarshal methods (with ptr receivers) are available too.
type Address = bytes.HexBytes

func AddressHash(bz []byte) Address {
	return Address(hash.SumTruncated(bz))
}
