package dependencies

import "context"

// InstallOptions contains options for dependency installation
type InstallOptions struct {
	SkipExisting bool   // Skip if already installed
	Verbose      bool   // Print verbose output
	Platform     string // windows or darwin
}

// InstallResult contains the result of an installation attempt
type InstallResult struct {
	Installed []string         // Tools that were successfully installed
	Failed    map[string]error // Tools that failed to install
	Skipped   []string         // Tools that were skipped
}

// DependenciesManager defines the interface for dependency management
type DependenciesManager interface {
	// CheckAndInstall verifies and installs missing dependencies
	CheckAndInstall(ctx context.Context, opts InstallOptions) (*InstallResult, error)

	// IsInstalled checks if a specific tool is installed
	IsInstalled(tool string) (bool, error)

	// DetectPackageManager returns the available package manager
	DetectPackageManager() (string, error)

	// GetRequiredTools returns the list of required tools
	GetRequiredTools() []Tool
}
