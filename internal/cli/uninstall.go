package cli

import (
	"context"
	"fmt"

	"github.com/jodify/jodify-setup/internal/installer"
	"github.com/jodify/jodify-setup/internal/platform"
	"github.com/spf13/cobra"
)

var (
	uninstallRestoreBackup bool
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall Jodify Neovim configuration",
	Long: `Remove the Jodify Neovim configuration from your system.

This command will:
1. Check if a previous installation exists
2. Remove the configuration directory
3. Optionally restore from backup if available

Use --restore-backup to restore a previous configuration.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		p, err := platform.Detect()
		if err != nil {
			return fmt.Errorf("failed to detect platform: %w", err)
		}

		opts := installer.UninstallOptions{
			RestoreBackup: uninstallRestoreBackup,
		}

		inst := installer.New(p)
		if err := inst.Uninstall(context.Background(), opts); err != nil {
			return fmt.Errorf("uninstall failed: %w", err)
		}

		fmt.Println("Jodify Neovim configuration uninstalled successfully!")
		return nil
	},
}

func init() {
	uninstallCmd.Flags().BoolVar(&uninstallRestoreBackup, "restore-backup", false, "Restore from backup if available")
}
