package file

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/crypto-provider/pkg/components"
	"github.com/cosmos/crypto-provider/pkg/factory"
)

// TODO: Should each provider have its own go.mod?

const (
	SourceMetadata = "metadata"
)

type FileProviderFactory struct {
	components.BaseFactory
}

// Register into the global factory
func init() {
	f := factory.GetGlobalFactory()
	err := f.RegisterFactory(&FileProviderFactory{})
	if err != nil {
		// TODO err instead of panic
		panic(fmt.Sprintf("failed to register factory: %v", err))
	}
}

var _ components.CryptoProviderFactory = (*FileProviderFactory)(nil)

func (f FileProviderFactory) Create(source components.BuildSource) (components.CryptoProvider, error) {
	switch s := source.(type) {
	case components.BuildSourceNew:
		return createDefault(s.Name)
	case components.BuildSourceMetadata:
		return createFromMetadata(s.Metadata)
	case components.BuildSourceJson:
		return createFromJson(s.JsonString)
	default:
		return nil, fmt.Errorf("unsupported source type: %T", source)
	}
}

func createDefault(name string) (*FileProvider, error) {
	meta := components.ProviderMetadata{
		Version: Version,
		Type:    ProviderTypeFile,
		Name:    name,
	}

	return createFromMetadata(meta)
}

func createFromMetadata(metadata components.ProviderMetadata) (*FileProvider, error) {
	if err := metadata.Validate(); err != nil {
		return nil, err
	}

	providerConfig, err := BuildConfig(metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to build config: %w", err)
	}

	if err := providerConfig.Validate(); err != nil {
		return nil, err
	}

	return &FileProvider{
		filePath: providerConfig.FilePath,
		metadata: metadata,
	}, nil
}

func createFromJson(jsonString string) (*FileProvider, error) {
	var metadata components.ProviderMetadata
	err := json.Unmarshal([]byte(jsonString), &metadata)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}
	return createFromMetadata(metadata)
}

func (FileProviderFactory) Type() string {
	return ProviderTypeFile
}

func (f FileProviderFactory) SupportedSources() []string {
	return []string{SourceMetadata, "new", "mnemonic", "json"}
}

// Add this method to implement the full interface
func (f FileProviderFactory) Save(cp components.CryptoProvider) error {
	return f.BaseFactory.Save(cp)
}
