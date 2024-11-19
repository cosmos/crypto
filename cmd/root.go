package main

import (
	"fmt"
	"github.com/cosmos/crypto-provider/pkg/cli"
	"github.com/cosmos/crypto-provider/pkg/keyring"
	"github.com/cosmos/crypto-provider/pkg/wallet"
	"os"

	"github.com/spf13/cobra"
)

var (
	flags struct {
		providersDir string
	}
)

// SimpleAddressFormatter implementation
type SimpleAddressFormatter struct{}

func (f SimpleAddressFormatter) FormatAddress(pubKey []byte) (string, error) {
	return fmt.Sprintf("addr_%x", pubKey[:8]), nil
}

func setup() (wallet.Wallet, error) {
	addressFormatter := SimpleAddressFormatter{}
	w, err := wallet.NewKeyringWallet("wallet-app", keyring.BackendMemory, flags.providersDir, addressFormatter)
	if err != nil {
		return nil, fmt.Errorf("failed to create wallet: %v", err)
	}
	return w, nil
}

func initFlags(rootCmd *cobra.Command) {
	rootCmd.PersistentFlags().StringVar(&flags.providersDir, "providers-dir", "", "Directory containing provider configurations")
	_ = rootCmd.MarkPersistentFlagRequired("providers-dir")
}

func main() {
	rootCmd := cli.GetRootCmd()
	initFlags(rootCmd)
	rootCmd.AddCommand(listCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
