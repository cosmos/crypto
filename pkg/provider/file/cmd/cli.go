package cmd

import (
	"crypto-provider/pkg/cli"
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	cmd := NewCommand()
	cli.GetRootCmd().AddCommand(cmd)
}

// NewCommand creates and returns the command for the file provider
func NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "file",
		Short: "Manage file-based crypto providers",
		Long:  `This command allows you to interact with file-based crypto providers.`,
	}

	// Add subcommands specific to file provider
	cmd.AddCommand(newCreateCommand())
	cmd.AddCommand(newListCommand())

	return cmd
}

func newCreateCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "create",
		Short: "Create a new file-based provider",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Creating a new file-based provider...")
			// Implement creation logic here
			return nil
		},
	}
}

func newListCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all file-based providers",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Listing all file-based providers...")
			// Implement listing logic here
			return nil
		},
	}
}
