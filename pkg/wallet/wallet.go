package wallet

import (
	"crypto-provider/cmd/register"
	"encoding/json"
	"fmt"
	"os"

	"crypto-provider/pkg/components"
	"crypto-provider/pkg/factory"
	"crypto-provider/pkg/keyring"
)

// Wallet interface defines the operations that can be performed on a wallet.
type Wallet interface {
	// NewCryptoProvider creates a new CryptoProvider and stores it in the Wallet
	NewCryptoProvider(providerType string, source components.BuildSource) error

	// GetCryptoProvider retrieves a CryptoProvider from the Wallet.
	GetCryptoProvider(id string) (components.CryptoProvider, error)

	// RetrieveCryptoProviderByAddress retrieves a CryptoProvider from the Wallet using its formatted address.
	RetrieveCryptoProviderByAddress(address string) (components.CryptoProvider, error)

	// ListProviders returns a list of all stored CryptoProvider UIDs.
	ListProviders() ([]string, error)

	// DeleteProvider removes a CryptoProvider from the Wallet.
	DeleteProvider(id string) error

	// GetProviderMetadata retrieves the metadata of a stored CryptoProvider.
	GetProviderMetadata(id string) (*components.ProviderMetadata, error)
}

// KeyringWallet implements the Wallet interface using a Keyring backend.
type KeyringWallet struct {
	kr               keyring.Keyring
	factory          *factory.Factory
	addressFormatter components.AddressFormatter
}

// NewKeyringWallet creates a new KeyringWallet with the specified parameters.
func NewKeyringWallet(appName, backend, rootDir string, addressFormatter components.AddressFormatter) (Wallet, error) {
	kr, err := keyring.NewKeyring(appName, backend, rootDir, os.Stdin)
	if err != nil {
		return nil, fmt.Errorf("failed to create keyring: %w", err)
	}

	// Init register
	register.Init()

	return &KeyringWallet{
		kr:               kr,
		factory:          factory.GetFactory(),
		addressFormatter: addressFormatter,
	}, nil
}

func (w *KeyringWallet) NewCryptoProvider(providerType string, source components.BuildSource) error {
	provider, err := w.factory.CreateCryptoProvider(providerType, source)
	if err != nil {
		return err
	}

	err = w.StoreCryptoProvider(provider.Metadata().Name, provider)
	if err != nil {
		return err
	}

	fmt.Printf("CryptoProvider of type '%s' created with PublicKey: %s\n", providerType, provider.Metadata().PublicKey)
	return nil
}

// StoreCryptoProvider stores a CryptoProvider in the Keyring.
func (w *KeyringWallet) StoreCryptoProvider(uid string, provider components.CryptoProvider) error {
	metadata := provider.Metadata()
	data, err := json.Marshal(metadata)
	if err != nil {
		return fmt.Errorf("failed to marshal CryptoProvider metadata: %w", err)
	}

	_, err = w.kr.NewItem(uid, data, "json")
	if err != nil {
		return fmt.Errorf("failed to store CryptoProvider: %w", err)
	}

	return nil
}

// GetCryptoProvider retrieves a CryptoProvider from the Keyring.
func (w *KeyringWallet) GetCryptoProvider(uid string) (components.CryptoProvider, error) {
	record, err := w.kr.Get(uid)
	if err != nil {
		return nil, fmt.Errorf("failed to get record: %w", err)
	}

	providerMeta, err := components.FromRecord(record)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal provider metadata: %w", err)
	}

	provider, err := w.factory.CreateCryptoProvider(providerMeta.Type, components.BuildSourceMetadata{Metadata: *providerMeta})
	if err != nil {
		return nil, fmt.Errorf("failed to create CryptoProvider from JSON: %w", err)
	}

	return provider, nil
}

// RetrieveCryptoProviderByAddress retrieves a CryptoProvider from the Wallet using its formatted address.
func (w *KeyringWallet) RetrieveCryptoProviderByAddress(address string) (components.CryptoProvider, error) {
	providers, err := w.ListProviders()
	if err != nil {
		return nil, fmt.Errorf("failed to list providers: %w", err)
	}

	for _, uid := range providers {
		metadata, err := w.GetProviderMetadata(uid)
		if err != nil {
			continue
		}

		formattedAddress, err := w.addressFormatter.FormatAddress([]byte(metadata.PublicKey))
		if err != nil {
			fmt.Printf("failed to format address: %v\n", err)
			continue
		}
		if formattedAddress == address {
			return w.GetCryptoProvider(uid)
		}
	}

	return nil, fmt.Errorf("no provider found with address: %s", address)
}

// ListProviders returns a list of all stored CryptoProvider UIDs.
func (w *KeyringWallet) ListProviders() ([]string, error) {
	records, err := w.kr.List()
	if err != nil {
		return nil, fmt.Errorf("failed to list providers: %w", err)
	}

	uids := make([]string, len(records))
	for i, record := range records {
		uids[i] = record.Key
	}

	return uids, nil
}

// DeleteProvider removes a CryptoProvider from the Keyring.
func (w *KeyringWallet) DeleteProvider(uid string) error {
	err := w.kr.Delete(uid)
	if err != nil {
		return fmt.Errorf("failed to delete provider: %w", err)
	}
	return nil
}

// GetProviderMetadata retrieves the metadata of a stored CryptoProvider.
func (w *KeyringWallet) GetProviderMetadata(uid string) (*components.ProviderMetadata, error) {
	record, err := w.kr.Get(uid)
	if err != nil {
		return nil, fmt.Errorf("failed to get provider metadata: %w", err)
	}

	return components.FromRecord(record)
}
