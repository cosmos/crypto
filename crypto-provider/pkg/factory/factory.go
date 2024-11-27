package factory

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/cosmos/crypto-provider/pkg/components"
)

var (
	factoryInstance *Factory
	once            sync.Once
)

type Factory struct {
	registry map[string]components.CryptoProviderFactory
}

func GetGlobalFactory() *Factory {
	once.Do(func() {
		factoryInstance = newFactory()
	})
	return factoryInstance
}

// GetRegisteredFactories returns a slice of all registered factory types.
func (f *Factory) GetRegisteredFactories() []string {
	factories := make([]string, 0, len(f.registry))
	for key := range f.registry {
		factories = append(factories, key)
	}
	return factories
}

// RegisterFactory is a function that registers a CryptoProviderFactory for its corresponding type.
func (f *Factory) RegisterFactory(factory components.CryptoProviderFactory) error {
	providerType := factory.Type()
	if providerType == "" {
		return fmt.Errorf("provider type cannot be empty")
	}

	if _, exists := f.registry[providerType]; exists {
		fmt.Printf("warning: factory for provider type '%s' already registered\n", providerType)
		return nil
	}

	f.registry[providerType] = factory
	return nil
}

// CreateCryptoProvider creates a CryptoProvider based on the provided metadata.
func (f *Factory) CreateCryptoProvider(providerType string, source components.BuildSource) (components.CryptoProvider, error) {
	factory, exists := f.registry[providerType]
	if !exists {
		return nil, fmt.Errorf("no factory registered for provider type: '%s'", providerType)
	}

	provider, err := factory.Create(source)
	if err != nil {
		return nil, err
	}

	// Type assert to internal interface
	initializer, ok := provider.(components.ProviderInitializer)
	if !ok {
		return nil, fmt.Errorf("provider does not implement initializer interface")
	}

	// Initialize keys using internal interface
	if err := initializer.InitializeKeys(); err != nil {
		return nil, fmt.Errorf("failed to initialize keys: %w. Check the implementation if InitializeKeys()", err)
	}

	// Try get pubkey to check if it was initialized
	pubKey := provider.GetPubKey()
	if pubKey == nil {
		return nil, fmt.Errorf("public key not available from provider")
	}

	return provider, nil
}

// LoadCryptoProvider loads a CryptoProvider from a raw JSON string.
func (f *Factory) LoadCryptoProvider(rawJSON string) (components.CryptoProvider, error) {
	var config components.CryptoProviderConfig
	if err := json.Unmarshal([]byte(rawJSON), &config); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	if config.ProviderType == "" {
		return nil, fmt.Errorf("provider_type is required in the configuration")
	}

	source := components.BuildSourceConfig{Config: config}
	return f.CreateCryptoProvider(config.ProviderType, source)
}

// newFactory creates a new Factory instance and initializes the registry map.
func newFactory() *Factory {
	return &Factory{
		registry: make(map[string]components.CryptoProviderFactory),
	}
}
