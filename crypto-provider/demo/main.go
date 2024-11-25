package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/cosmos/crypto-provider/pkg/components"
	"github.com/cosmos/crypto-provider/pkg/impl/file"
	"github.com/cosmos/crypto-provider/pkg/keyring"
	"github.com/cosmos/crypto-provider/pkg/wallet"
)

const TestFile = "testdata/file_1.json"

// SimpleAddressFormatter implementation
type SimpleAddressFormatter struct{}

func (f SimpleAddressFormatter) FormatAddress(pubKey []byte) (string, error) {
	return fmt.Sprintf("addr_%x", pubKey[:8]), nil
}

func main() {
	currentDir, _ := os.Getwd()
	providersDir := filepath.Join(currentDir, "testdata")

	// Step 1: Create a new wallet
	addressFormatter := SimpleAddressFormatter{}
	w, err := wallet.NewKeyringWallet("demo-app", keyring.BackendMemory, providersDir, addressFormatter)
	if err != nil {
		log.Fatalf("Failed to create wallet: %v", err)
	}

	// Step 2: Load JSON file and create a new crypto provider
	jsonFilePath := filepath.Join(currentDir, TestFile)
	jsonData, err := os.ReadFile(jsonFilePath)
	if err != nil {
		log.Fatalf("Failed to read JSON file: %v", err)
	}

	buildSource := components.BuildSourceJson{JsonString: string(jsonData)}
	err = w.NewCryptoProvider(file.ProviderTypeFile, buildSource)
	if err != nil {
		log.Fatalf("Failed to create new crypto provider: %v", err)
	}

	// Step 3: List providers and get the first one
	providerIDs, err := w.ListProviders()
	if err != nil {
		log.Fatalf("Failed to list providers: %v", err)
	}
	if len(providerIDs) == 0 {
		log.Fatalf("No providers found")
	}

	provider, err := w.GetCryptoProvider(providerIDs[0])
	if err != nil {
		log.Fatalf("Failed to get provider: %v", err)
	}

	// Step 4: Get signer from the provider
	signer := provider.GetSigner()

	// Step 5: Generate random data and sign it
	randomBytes := make([]byte, 32)
	_, err = rand.Read(randomBytes)
	if err != nil {
		log.Fatalf("Failed to generate random bytes: %v", err)
	}

	signature, err := signer.Sign(randomBytes, nil)
	if err != nil {
		log.Fatalf("Failed to sign data: %v", err)
	}

	// Step 6: Get verifier and verify the signature
	verifier := provider.GetVerifier()
	valid, err := verifier.Verify(signature, randomBytes, nil)
	if err != nil {
		log.Fatalf("Failed to verify signature: %v", err)
	}

	fmt.Printf("Signature verification result: %v\n", valid)
}
