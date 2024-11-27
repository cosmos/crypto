package keyring

import "github.com/99designs/keyring"

// Record a generic wrapper for keyring.Item
type Record struct {
	Key       string
	Data      []byte
	CodecType string
}

// NewRecord creates a new Record
func NewRecord(name string, data []byte, codecType string) *Record {
	return &Record{
		Key:       name,
		Data:      data,
		CodecType: codecType,
	}
}

// ToItem converts the Record to a keyring.Item
func (r *Record) ToItem() keyring.Item {
	return keyring.Item{
		Key:         r.Key,
		Data:        r.Data,
		Description: r.CodecType,
	}
}

// FromItem creates a Record from a keyring.Item
func FromItem(item keyring.Item) *Record {
	return &Record{
		Key:       item.Key,
		Data:      item.Data,
		CodecType: item.Description,
	}
}
