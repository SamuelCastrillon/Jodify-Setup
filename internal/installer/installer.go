package installer

import (
	"context"
	"fmt"
	"os"
	"runtime"

	"github.com/SamuelCastrillon/Jodify-Setup/internal/config"
	"github.com/SamuelCastrillon/Jodify-Setup/internal/dependencies"
	"github.com/SamuelCastrillon/Jodify-Setup/internal/platform"
)

// Installer orchestrates the setup process
type Installer interface {
	// Install performs a fresh installation
	Install(ctx context.Context, opts InstallOptions) error

	// Update refreshes the configuration
	Update(ctx context.Context, opts UpdateOptions) error

	// Uninstall removes the configuration
	Uninstall(ctx context.Context, opts UninstallOptions) error
}

// installer implements Installer
type installer struct {
	platform  platform.Platform
	configMgr config.Manager
}

// InstallOptions contains options for installation
type InstallOptions struct {
	Force      bool   // Overwrite existing config
	ConfigURL  string // Custom config URL (optional)
	SkipBackup bool   // Don't create backup
	SkipDeps   bool   // Skip installing dependencies
}

// UpdateOptions contains options for update
type UpdateOptions struct {
	ConfigURL string // Custom config URL (optional)
}

// UninstallOptions contains options for uninstall
type UninstallOptions struct {
	RestoreBackup bool // Restore from backup if available
}

// New creates a new Installer
func New(p platform.Platform) Installer {
	mgr := config.NewManager()
	return &installer{
		platform:  p,
		configMgr: mgr,
	}
}

// NewWithConfig creates a new Installer with injected config manager (for testing)
func NewWithConfig(p platform.Platform, mgr config.Manager) Installer {
	return &installer{
		platform:  p,
		configMgr: mgr,
	}
}

// Install performs a fresh installation
func (i *installer) Install(ctx context.Context, opts InstallOptions) error {
	// Check prerequisites first
	result, err := i.platform.CheckPrerequisites()
	if err != nil {
		return fmt.Errorf("failed to check prerequisites: %w", err)
	}

	if !result.NeovimInstalled {
		return fmt.Errorf("%w: please install Neovim first (https://github.com/neovim/neovim/wiki/Installing-Neovim)", platform.ErrNeovimNotFound)
	}

	if !result.GitInstalled {
		return fmt.Errorf("%w: please install Git first (https://git-scm.com/)", platform.ErrGitNotFound)
	}

	// Install dependencies if not skipped
	if !opts.SkipDeps {
		var depsMgr dependencies.DependenciesManager
		switch runtime.GOOS {
		case "windows":
			depsMgr = dependencies.NewWindowsDeps()
		case "darwin":
			depsMgr = dependencies.NewDarwinDeps()
		default:
			fmt.Printf("Warning: unsupported platform for dependencies: %s\n", runtime.GOOS)
		}

		if depsMgr != nil {
			fmt.Println("Checking and installing required tools...")
			depsResult, err := depsMgr.CheckAndInstall(ctx, dependencies.InstallOptions{
				Verbose: true,
			})
			if err != nil {
				fmt.Printf("Warning: dependency check failed: %v\n", err)
			}

			if len(depsResult.Installed) > 0 {
				fmt.Printf("Installed tools: %v\n", depsResult.Installed)
			}
			if len(depsResult.Failed) > 0 {
				fmt.Printf("Failed tools: %v\n", depsResult.Failed)
			}
		}
	}

	// Get config directory
	configDir, err := i.platform.GetConfigDir()
	if err != nil {
		return fmt.Errorf("failed to get config directory: %w", err)
	}

	// Check if existing config exists
	if !opts.Force && !opts.SkipBackup {
		if _, err := os.Stat(configDir); err == nil {
			// Config exists, create backup
			backupDir := configDir + ".backup"
			if err := os.Rename(configDir, backupDir); err != nil {
				return fmt.Errorf("failed to create backup: %w", err)
			}
			fmt.Printf("Existing config backed up to: %s\n", backupDir)
		}
	}

	// Create config directory
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	// Determine download URL
	downloadURL := opts.ConfigURL
	if downloadURL == "" {
		downloadURL = i.configMgr.GetLatestReleaseURL()
	}

	// Download config
	zipPath, err := i.configMgr.Download(ctx, downloadURL)
	if err != nil {
		return fmt.Errorf("failed to download config: %w", err)
	}
	defer os.Remove(zipPath)

	// Extract config
	if err := i.configMgr.Extract(zipPath, configDir); err != nil {
		return fmt.Errorf("failed to extract config: %w", err)
	}

	return nil
}

// Update refreshes the configuration
func (i *installer) Update(ctx context.Context, opts UpdateOptions) error {
	configDir, err := i.platform.GetConfigDir()
	if err != nil {
		return fmt.Errorf("failed to get config directory: %w", err)
	}

	// Check if config exists
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		return fmt.Errorf("no previous installation found: run 'jodify-setup install' first")
	}

	// Determine download URL
	downloadURL := opts.ConfigURL
	if downloadURL == "" {
		downloadURL = i.configMgr.GetLatestReleaseURL()
	}

	// Download config
	zipPath, err := i.configMgr.Download(ctx, downloadURL)
	if err != nil {
		return fmt.Errorf("failed to download config: %w", err)
	}
	defer os.Remove(zipPath)

	// Extract config (replace existing)
	if err := i.configMgr.Extract(zipPath, configDir); err != nil {
		return fmt.Errorf("failed to extract config: %w", err)
	}

	return nil
}

// Uninstall removes the configuration
func (i *installer) Uninstall(ctx context.Context, opts UninstallOptions) error {
	configDir, err := i.platform.GetConfigDir()
	if err != nil {
		return fmt.Errorf("failed to get config directory: %w", err)
	}

	// Check if config exists
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		return nil // Nothing to uninstall
	}

	// Check for backup
	backupDir := configDir + ".backup"
	hasBackup := false
	if _, err := os.Stat(backupDir); err == nil {
		hasBackup = true
	}

	// Remove config directory
	if err := os.RemoveAll(configDir); err != nil {
		return fmt.Errorf("failed to remove config: %w", err)
	}

	// Restore backup if requested and available
	if opts.RestoreBackup && hasBackup {
		if err := os.Rename(backupDir, configDir); err != nil {
			return fmt.Errorf("failed to restore backup: %w", err)
		}
		fmt.Printf("Restored from backup: %s\n", backupDir)
	}

	return nil
}

// CheckPrerequisites validates the system is ready
func (i *installer) CheckPrerequisites() error {
	result, err := i.platform.CheckPrerequisites()
	if err != nil {
		return err
	}

	if len(result.Errors) > 0 {
		return result.Errors[0]
	}

	return nil
}

// GetConfigDir returns the config directory for the platform
func (i *installer) GetConfigDir() (string, error) {
	return i.platform.GetConfigDir()
}

// GetConfigBackupDir returns the backup directory path
func GetConfigBackupDir(configDir string) string {
	return configDir + ".backup"
}

// ConfigExists checks if a config exists at the given path
func ConfigExists(configDir string) (bool, error) {
	_, err := os.Stat(configDir)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// BackupConfig creates a backup of the existing config
func BackupConfig(configDir string) (string, error) {
	backupDir := GetConfigBackupDir(configDir)
	if err := os.Rename(configDir, backupDir); err != nil {
		return "", fmt.Errorf("failed to backup: %w", err)
	}
	return backupDir, nil
}

// RestoreBackup restores config from backup
func RestoreBackup(configDir string) error {
	backupDir := GetConfigBackupDir(configDir)
	if err := os.Rename(backupDir, configDir); err != nil {
		return fmt.Errorf("failed to restore backup: %w", err)
	}
	return nil
}
