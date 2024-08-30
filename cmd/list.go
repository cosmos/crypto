package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all providers",
	Long:  `This command lists all the crypto providers currently available in the wallet.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		w, err := setup()
		if err != nil {
			return err
		}

		providers, err := w.ListProviders()
		if err != nil {
			return fmt.Errorf("failed to list providers: %v", err)
		}

		if len(providers) == 0 {
			fmt.Println("No providers found.")
		} else {
			fmt.Println("Available providers:")
			for _, provider := range providers {
				fmt.Printf("- %s\n", provider)
			}
		}

		return nil
	},
}
