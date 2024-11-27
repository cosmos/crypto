package cli

import (
	"github.com/spf13/cobra"
	"sync"
)

var (
	rootCmd     *cobra.Command
	initRootCmd sync.Once
)

// GetRootCmd returns the root command for the application
func GetRootCmd() *cobra.Command {
	initRootCmd.Do(func() {
		rootCmd = &cobra.Command{
			Use:   "wallet",
			Short: "Wallet is a CLI for managing crypto providers",
			Long:  `Wallet is a command-line application for managing and interacting with various crypto providers.`,
		}
	})

	return rootCmd
}
