package cli

import (
	"fmt"
	"os"

	"github.com/SamuelCastrillon/Jodify-Setup/internal/platform"
	"github.com/spf13/cobra"
)

var (
	cleanKeepBackup bool
)

var devCleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean Jodify config and optionally backup",
	Long: `Clean the Jodify Neovim configuration from your system.

This command removes the Jodify config directory. Use --keep-backup
to preserve any existing backup.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		p, err := platform.Detect()
		if err != nil {
			return fmt.Errorf("failed to detect platform: %w", err)
		}

		configDir, err := p.GetConfigDir()
		if err != nil {
			return fmt.Errorf("failed to get config directory: %w", err)
		}

		// Check if config exists
		exists, err := configExists(configDir)
		if err != nil {
			return fmt.Errorf("failed to check config: %w", err)
		}

		if !exists {
			fmt.Println("No Jodify config found. Nothing to clean.")
			return nil
		}

		// Check for backup
		backupDir := configDir + ".backup"
		backupExists, _ := configExists(backupDir)

		// Remove config
		fmt.Printf("Removing config: %s\n", configDir)
		if err := os.RemoveAll(configDir); err != nil {
			return fmt.Errorf("failed to remove config: %w", err)
		}

		// Handle backup
		if backupExists && !cleanKeepBackup {
			fmt.Printf("Removing backup: %s\n", backupDir)
			if err := os.RemoveAll(backupDir); err != nil {
				fmt.Printf("Warning: failed to remove backup: %v\n", err)
			}
		} else if backupExists && cleanKeepBackup {
			fmt.Printf("Backup preserved: %s\n", backupDir)
		}

		fmt.Println("Clean completed!")
		return nil
	},
}

func configExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func init() {
	devCleanCmd.Flags().BoolVar(&cleanKeepBackup, "keep-backup", false, "Keep backup directory")
}
