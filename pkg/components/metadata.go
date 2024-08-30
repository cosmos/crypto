package components

import (
	"encoding/json"
	"fmt"

	"github.com/Masterminds/semver/v3"

	"crypto-provider/pkg/keyring"
)

type ProviderConfig = map[string]any

type ProviderMetadata struct {
	Version   string         `json:"version"`
	Name      string         `json:"name"`
	Type      string         `json:"type"`
	PublicKey string         `json:"publickey"`
	Config    ProviderConfig `json:"config"`
}

func FromRecord(record *keyring.Record) (*ProviderMetadata, error) {
	var meta ProviderMetadata
	if record.CodecType != "json" {
		return nil, fmt.Errorf("unsupported codec type: %s", record.CodecType)
	}

	err := json.Unmarshal(record.Data, &meta)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal metadata: %w", err)
	}

	return &meta, nil
}

// Validate checks if the ProviderMetadata is valid
func (pm *ProviderMetadata) Validate() error {
	_, err := semver.NewVersion(pm.Version)
	if err != nil {
		return fmt.Errorf("invalid version: %w", err)
	}

	if pm.Name == "" {
		return fmt.Errorf("name is required")
	}
	if pm.Type == "" {
		return fmt.Errorf("type is required")
	}

	if pm.PublicKey == "" {
		return fmt.Errorf("publickey is required")
	}

	return nil
}
