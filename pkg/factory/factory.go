package factory

import (
	"crypto-provider/pkg/components"
	"fmt"
	"sync"
)

var (
	factoryInstance *Factory
	once            sync.Once
)

type Factory struct {
	registry map[string]components.CryptoProviderFactory
}

func GetFactory() *Factory {
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
		return fmt.Errorf("provider type %s is already registered", providerType)
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

	return factory.Create(source)
}

// newFactory creates a new Factory instance and initializes the registry map.
func newFactory() *Factory {
	return &Factory{
		registry: make(map[string]components.CryptoProviderFactory),
	}
}
