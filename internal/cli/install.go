package cli

import (
	"context"
	"fmt"

	"github.com/SamuelCastrillon/Jodify-Setup/internal/installer"
	"github.com/SamuelCastrillon/Jodify-Setup/internal/platform"
	"github.com/spf13/cobra"
)

var (
	installForce     bool
	installSkipBkp   bool
	installConfigURL string
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install Jodify Neovim configuration",
	Long: `Install the Jodify Neovim configuration to your system.

This command will:
1. Check for Neovim and Git prerequisites
2. Create a backup of existing configuration if present
3. Download the latest configuration package
4. Extract it to the appropriate Neovim config directory

Use --force to overwrite existing configuration without prompting.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		p, err := platform.Detect()
		if err != nil {
			return fmt.Errorf("failed to detect platform: %w", err)
		}

		opts := installer.InstallOptions{
			Force:      installForce,
			ConfigURL:  installConfigURL,
			SkipBackup: installSkipBkp,
		}

		inst := installer.New(p)
		if err := inst.Install(context.Background(), opts); err != nil {
			return fmt.Errorf("installation failed: %w", err)
		}

		fmt.Println("Jodify Neovim configuration installed successfully!")
		return nil
	},
}

func init() {
	installCmd.Flags().BoolVarP(&installForce, "force", "f", false, "Overwrite existing configuration")
	installCmd.Flags().BoolVar(&installSkipBkp, "skip-backup", false, "Skip creating backup of existing configuration")
	installCmd.Flags().StringVar(&installConfigURL, "config-url", "", "Custom configuration URL (optional)")
}
