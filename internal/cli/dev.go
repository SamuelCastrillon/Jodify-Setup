package cli

import (
	"os"

	"github.com/spf13/cobra"
)

// devCmd is the parent command for development tools
var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "Development commands",
	Long: `Development commands for testing and debugging Jodify.

These commands help with local development and testing:
  sync     - Sync local config to Neovim
  open     - Open Neovim with Jodify config
  clean    - Clean Jodify config

Examples:
  jodify-setup sync --backup
  jodify-setup dev open
  jodify-setup dev clean`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	// Add dev command
	RootCmd.AddCommand(devCmd)

	// Add subcommands to dev
	devCmd.AddCommand(syncCmd)
	devCmd.AddCommand(devOpenCmd)
	devCmd.AddCommand(devCleanCmd)

	// Set output to stdout
	devCmd.SetOutput(os.Stdout)
}
