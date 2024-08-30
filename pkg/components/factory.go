package components

import (
	"encoding/json"
)

// CryptoProviderFactory is a factory interface for creating CryptoProviders.
// Must be implemented by each CryptoProvider.
type CryptoProviderFactory interface {
	// Create creates a new CryptoProvider instance using the given source
	Create(source BuildSource) (CryptoProvider, error)
	// Type returns the type of the CryptoProvider that this factory creates
	Type() string
	// SupportedSources returns the sources that this factory supports building CryptoProviders from
	SupportedSources() []string
}

type BuildSource interface {
	Type() string
	Validate() error
}

// Common BuildSource implementations

// BuildSourceNew /////////////////////////////////////////////////////////////////
// is a BuildSource implementation that creates a new provider with default values
// /////////////////////////////////////////////////////////////////////////////////
type BuildSourceNew struct {
	Name string
}

func (m BuildSourceNew) Type() string    { return "new" }
func (m BuildSourceNew) Validate() error { return nil }

// BuildSourceMetadata /////////////////////////////////////////////////////////////
// is a BuildSource implementation that uses ProviderMetadata as source
// /////////////////////////////////////////////////////////////////////////////////
type BuildSourceMetadata struct {
	Metadata ProviderMetadata
}

func (m BuildSourceMetadata) Type() string    { return "metadata" }
func (m BuildSourceMetadata) Validate() error { return m.Metadata.Validate() }

// BuildSourceMnemonic /////////////////////////////////////////////////////////////
// is a BuildSource implementation that uses a mnemonic as source
// /////////////////////////////////////////////////////////////////////////////////
type BuildSourceMnemonic struct {
	Mnemonic string
}

func (m BuildSourceMnemonic) Type() string { return "mnemonic" }
func (m BuildSourceMnemonic) Validate() error {
	//TODO
	return nil
}

// BuildSourceJson /////////////////////////////////////////////////////////////////
// is a BuildSource implementation that uses a JSON string as source
// /////////////////////////////////////////////////////////////////////////////////
type BuildSourceJson struct {
	JsonString string
}

func (m BuildSourceJson) Type() string { return "json" }
func (m BuildSourceJson) Validate() error {
	var jsonData map[string]interface{}
	err := json.Unmarshal([]byte(m.JsonString), &jsonData)
	if err != nil {
		return err
	}
	// Additional validation can be added here if needed
	return nil
}
