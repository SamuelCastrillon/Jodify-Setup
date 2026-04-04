package cli

import (
	"context"
	"fmt"

	"github.com/jodify/jodify-setup/internal/installer"
	"github.com/jodify/jodify-setup/internal/platform"
	"github.com/spf13/cobra"
)

var (
	updateConfigURL string
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Jodify Neovim configuration",
	Long: `Update the Jodify Neovim configuration to the latest version.

This command will:
1. Check if a previous installation exists
2. Download the latest configuration package
3. Extract it to replace the existing configuration

Use --config-url to update from a custom URL.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		p, err := platform.Detect()
		if err != nil {
			return fmt.Errorf("failed to detect platform: %w", err)
		}

		opts := installer.UpdateOptions{
			ConfigURL: updateConfigURL,
		}

		inst := installer.New(p)
		if err := inst.Update(context.Background(), opts); err != nil {
			return fmt.Errorf("update failed: %w", err)
		}

		fmt.Println("Jodify Neovim configuration updated successfully!")
		return nil
	},
}

func init() {
	updateCmd.Flags().StringVar(&updateConfigURL, "config-url", "", "Custom configuration URL (optional)")
}
