package file

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/crypto-provider/pkg/components"
)

// FileProviderConfig holds the configuration for the File Provider
type FileProviderConfig struct {
	FilePath string `json:"filepath"`
}

// BuildConfig creates a FileProviderConfig from the provided metadata
func BuildConfig(metadata components.ProviderMetadata) (FileProviderConfig, error) {
	var config FileProviderConfig
	jsonData, err := json.Marshal(metadata.Config)
	if err != nil {
		return FileProviderConfig{}, fmt.Errorf("failed to marshal config: %w", err)
	}

	err = json.Unmarshal(jsonData, &config)
	if err != nil {
		return FileProviderConfig{}, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return config, nil
}

// Validate checks if the FileProviderConfig is valid
func (c FileProviderConfig) Validate() error {
	if c.FilePath == "" {
		return fmt.Errorf("FilePath cannot be empty")
	}
	return nil
}
