package cli

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
)

// ErrCheckFlag is returned when --check flag is used
var ErrCheckFlag = errors.New("check flag handled")

// RootCmd represents the root command
var RootCmd = &cobra.Command{
	Use:   "jodify-setup",
	Short: "Jodify-Setup - Neovim configuration manager",
	Long: `Jodify-Setup is a CLI tool for managing Jodify Neovim configuration.

It provides commands to install, update, and uninstall the Jodify
Neovim configuration with full support for NVIM_APPNAME isolation.

Usage:
  jodify-setup install   Install Jodify Neovim configuration
  jodify-setup update   Update existing configuration
  jodify-setup uninstall Remove installed configuration
  jodify-setup version  Show version information`,
	Run: func(cmd *cobra.Command, args []string) {
		// Default behavior shows help
		cmd.Help()
	},
}

// Execute runs all commands
func Execute() error {
	return RootCmd.Execute()
}

func init() {
	// Add version command
	RootCmd.AddCommand(versionCmd)

	// Add install command
	RootCmd.AddCommand(installCmd)

	// Add update command
	RootCmd.AddCommand(updateCmd)

	// Add uninstall command
	RootCmd.AddCommand(uninstallCmd)

	// Set output to stdout
	RootCmd.SetOutput(os.Stdout)
}
