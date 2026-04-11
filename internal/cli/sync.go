package cli

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/SamuelCastrillon/Jodify-Setup/internal/platform"
	"github.com/spf13/cobra"
)

var (
	syncDryRun    bool
	syncBackup    bool
	syncSourceDir string
	syncProfile   string
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync local config to Neovim",
	Long: `Sync the local Jodify configuration to your Neovim config directory.

This command copies the local 'config/' directory to your Neovim config
directory for local development and testing.

By default, syncs to 'jodify' profile. Use --profile to sync to a different profile.

Examples:
  jodify-setup dev sync                    # Sync to jodify profile
  jodify-setup dev sync --profile nvim     # Sync to default nvim
  jodify-setup dev sync --dry-run          # Preview without copying
  jodify-setup dev sync --backup           # Backup before syncing`,
	RunE: func(cmd *cobra.Command, args []string) error {
		p, err := platform.Detect()
		if err != nil {
			return fmt.Errorf("failed to detect platform: %w", err)
		}

		// Set NVIM_APPNAME if profile is specified
		if syncProfile != "" {
			os.Setenv("NVIM_APPNAME", syncProfile)
			fmt.Printf("Using profile: %s\n", syncProfile)
		}

		// Determine source directory
		sourceDir := syncSourceDir
		if sourceDir == "" {
			// Default to ./config relative to current working directory
			cwd, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("failed to get current directory: %w", err)
			}
			sourceDir = filepath.Join(cwd, "config")
		}

		// Check source directory exists
		if _, err := os.Stat(sourceDir); os.IsNotExist(err) {
			return fmt.Errorf("source directory does not exist: %s", sourceDir)
		}

		// Get target Neovim config directory (uses NVIM_APPNAME env or default)
		targetDir, err := p.GetConfigDir()
		if err != nil {
			return fmt.Errorf("failed to get config directory: %w", err)
		}

		fmt.Printf("Source: %s\n", sourceDir)
		fmt.Printf("Target: %s\n", targetDir)

		// Check if target exists
		if _, err := os.Stat(targetDir); err == nil && !syncDryRun {
			if syncBackup {
				// Create backup
				backupDir := targetDir + ".backup"
				fmt.Printf("Creating backup: %s\n", backupDir)
				if err := os.Rename(targetDir, backupDir); err != nil {
					return fmt.Errorf("failed to create backup: %w", err)
				}
			} else {
				// Just remove existing
				fmt.Printf("Removing existing config: %s\n", targetDir)
				if err := os.RemoveAll(targetDir); err != nil {
					return fmt.Errorf("failed to remove existing config: %w", err)
				}
			}
		}

		if syncDryRun {
			fmt.Println("\n[DRY RUN] No files were copied")
			fmt.Println("\nFiles that would be copied:")
			return listFiles(sourceDir, "")
		}

		// Copy files
		fmt.Println("\nSyncing files...")
		if err := copyDir(sourceDir, targetDir); err != nil {
			return fmt.Errorf("failed to sync: %w", err)
		}

		fmt.Println("Config synced successfully!")
		if syncProfile != "" {
			fmt.Printf("\nTo test: NVIM_APPNAME=%s nvim\n", syncProfile)
		} else {
			fmt.Printf("\nTo test: NVIM_APPNAME=jodify nvim\n")
		}

		return nil
	},
}

// listFiles lists all files in the source directory
func listFiles(root, indent string) error {
	entries, err := os.ReadDir(root)
	if err != nil {
		return err
	}

	for _, e := range entries {
		name := e.Name()
		if e.IsDir() {
			fmt.Printf("%s%s/\n", indent, name)
			if err := listFiles(filepath.Join(root, name), indent+"  "); err != nil {
				return err
			}
		} else {
			fmt.Printf("%s%s\n", indent, name)
		}
	}
	return nil
}

// copyDir recursively copies a directory
func copyDir(src, dst string) error {
	// Create destination directory
	if err := os.MkdirAll(dst, 0755); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	entries, err := os.ReadDir(src)
	if err != nil {
		return err
	}

	for _, e := range entries {
		srcPath := filepath.Join(src, e.Name())
		dstPath := filepath.Join(dst, e.Name())

		if e.IsDir() {
			if err := copyDir(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			// Copy file
			data, err := os.ReadFile(srcPath)
			if err != nil {
				return err
			}
			if err := os.WriteFile(dstPath, data, e.Type()); err != nil {
				return err
			}
		}
	}

	return nil
}

func init() {
	syncCmd.Flags().BoolVar(&syncDryRun, "dry-run", false, "Preview changes without copying")
	syncCmd.Flags().BoolVar(&syncBackup, "backup", false, "Create backup before syncing")
	syncCmd.Flags().StringVar(&syncSourceDir, "source", "", "Source directory (default: ./config)")
	syncCmd.Flags().StringVarP(&syncProfile, "profile", "p", "jodify", "Neovim profile (NVIM_APPNAME)")
}
