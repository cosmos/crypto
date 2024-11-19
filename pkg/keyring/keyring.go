package keyring

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/99designs/keyring"
	
	"github.com/cosmos/crypto-provider/pkg/keyring/internal"
)

// Backend options for Keyring
const (
	BackendFile   = "file"
	BackendOS     = "os"
	BackendTest   = "test"
	BackendMemory = "memory"
)

//nolint:unused
const (
	keyringFileDirName = "keyring-file"
	keyringTestDirName = "keyring-test"
	passKeyringPrefix  = "keyring-%s"
)

var (
	_ Keyring = &keystore{}
)

// Keyring exposes operations over a backend supported by github.com/99designs/keyring.
type Keyring interface {
	// List returns all records in the keyring.
	List() ([]*Record, error)

	// Get retrieves a record from the keyring by its uid.
	Get(uid string) (*Record, error)

	// Delete removes a record from the keyring by its uid.
	Delete(uid string) error

	// NewItem creates a new item and stores it in the keyring.
	NewItem(uid string, data []byte, codecType string) (*Record, error)
}

// NewKeyring creates a new instance of a keyring with the specified backend.
func NewKeyring(appName, backend, rootDir string, userInput io.Reader) (Keyring, error) {
	var (
		db  keyring.Keyring
		err error
	)

	switch backend {
	case BackendMemory:
		return newInMemoryWithKeyring(keyring.NewArrayKeyring(nil)), err
	case BackendTest:
		db, err = keyring.Open(newTestBackendKeyringConfig(appName, rootDir))
	case BackendFile:
		var cfg keyring.Config
		cfg, err = newFileBackendKeyringConfig(appName, rootDir, userInput)
		if err != nil {
			return nil, err
		}
		db, err = keyring.Open(cfg)
	case BackendOS:
		db, err = keyring.Open(newOSBackendKeyringConfig(appName, rootDir, userInput))
	default:
		return nil, fmt.Errorf("no available implementation for backend: %s", backend)
	}

	if err != nil {
		return nil, err
	}

	return newKeystore(db, backend), nil
}

// newInMemoryWithKeyring returns an in-memory keyring using the specified keyring.Keyring as the backing store.
func newInMemoryWithKeyring(kr keyring.Keyring) Keyring {
	return newKeystore(kr, BackendMemory)
}

type keystore struct {
	db      keyring.Keyring
	backend string
}

// List returns all records in the keystore.
func (ks keystore) List() ([]*Record, error) {
	items, err := ks.db.Keys()
	if err != nil {
		return nil, err
	}

	var records []*Record
	for _, key := range items {
		item, err := ks.db.Get(key)
		if err != nil {
			return nil, err
		}
		records = append(records, FromItem(item))
	}

	return records, nil
}

// NewItem creates a new item and stores it in the keystore.
func (ks keystore) NewItem(uid string, data []byte, codecType string) (*Record, error) {
	record := NewRecord(uid, data, codecType)
	item := record.ToItem()

	err := ks.db.Set(item)
	if err != nil {
		return nil, err
	}

	return record, nil
}

// Get retrieves a record from the keystore by its uid.
func (ks keystore) Get(uid string) (*Record, error) {
	item, err := ks.db.Get(uid)
	if err != nil {
		if errors.Is(err, keyring.ErrKeyNotFound) {
			return nil, keyring.ErrKeyNotFound
		}
		return nil, err
	}

	return FromItem(item), nil
}

// Delete removes a record from the keystore by its uid.
func (ks keystore) Delete(uid string) error {
	err := ks.db.Remove(uid)
	if err != nil {
		if errors.Is(err, keyring.ErrKeyNotFound) {
			return keyring.ErrKeyNotFound
		}
		return err
	}
	return nil
}

// newKeystore creates a new keystore instance with the given keyring and backend.
func newKeystore(kr keyring.Keyring, backend string) keystore {
	return keystore{
		db:      kr,
		backend: backend,
	}
}

// newOSBackendKeyringConfig creates a new OS backend keyring configuration.
func newOSBackendKeyringConfig(appName, dir string, buf io.Reader) keyring.Config {
	return keyring.Config{
		ServiceName:              appName,
		FileDir:                  dir,
		KeychainTrustApplication: true,
		FilePasswordFunc:         internal.NewRealPrompt(dir, buf),
	}
}

// newTestBackendKeyringConfig creates a new test backend keyring configuration.
func newTestBackendKeyringConfig(appName, dir string) keyring.Config {
	return keyring.Config{
		AllowedBackends: []keyring.BackendType{keyring.FileBackend},
		ServiceName:     appName,
		FileDir:         dir,
		FilePasswordFunc: func(_ string) (string, error) {
			return "test", nil
		},
	}
}

// newPassBackendKeyringConfig creates a new pass backend keyring configuration.
//
//nolint:unused
func newPassBackendKeyringConfig(appName, _ string, _ io.Reader) (keyring.Config, error) {
	prefix := fmt.Sprintf(passKeyringPrefix, appName)

	return keyring.Config{
		AllowedBackends: []keyring.BackendType{keyring.PassBackend},
		ServiceName:     appName,
		PassPrefix:      prefix,
	}, nil
}

// newFileBackendKeyringConfig creates a new file backend keyring configuration.
func newFileBackendKeyringConfig(name, dir string, buf io.Reader) (keyring.Config, error) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0700)
		if err != nil {
			return keyring.Config{}, fmt.Errorf("failed to create keyring directory: %w", err)
		}
	}

	return keyring.Config{
		AllowedBackends:  []keyring.BackendType{keyring.FileBackend},
		ServiceName:      name,
		FileDir:          dir,
		FilePasswordFunc: internal.NewRealPrompt(dir, buf),
	}, nil
}
