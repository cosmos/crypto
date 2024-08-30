/*
Package cli implements a "pluggable" command-line interface (CLI) system for managing specific CLI commands of each provider type in a way that doesn't
require modifying the core code

Key components:

1. Root Command (root.go):
   - Implements a GetRootCmd() function using the singleton pattern and sync.Once for thread-safety.
   - Initializes the root command and its flags.

2. Provider Commands (e.g., pkg/provider/file/cmd/command.go):
   - Each provider package implements its own set of subcommands.
   - Uses an init() function to register its commands with the root command.

How to add a new provider cli:

1. Create a new package for your provider (e.g., pkg/provider/newprovider/cmd/).
2. In this package, create a file (e.g., command.go) with the following structure:
   - Implement an init() function that calls GetRootCmd() and adds your provider's command.
   - Create a NewCommand() function that returns a cobra.Command for your provider.
   - Implement any subcommands specific to your provider.

3. Import your new provider package in the register.go file for side effects.
*/

package cli
