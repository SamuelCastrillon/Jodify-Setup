package dependencies

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
)

// windowsDeps implements DependenciesManager for Windows using Scoop
type windowsDeps struct{}

// NewWindowsDeps creates a new Windows dependencies manager
func NewWindowsDeps() DependenciesManager {
	return &windowsDeps{}
}

// DetectPackageManager returns the available package manager on Windows
func (w *windowsDeps) DetectPackageManager() (string, error) {
	// Check if scoop is available
	cmd := exec.Command("scoop", "--version")
	if err := cmd.Run(); err == nil {
		return "scoop", nil
	}
	return "", fmt.Errorf("no package manager found (scoop not available)")
}

// IsInstalled checks if a tool is installed on Windows
func (w *windowsDeps) IsInstalled(tool string) (bool, error) {
	// Try to get version of the tool
	cmd := exec.Command(tool, "--version")
	output, err := cmd.Output()
	if err != nil {
		// Try alternative: scoop list tool
		cmd = exec.Command("scoop", "list", tool)
		if err := cmd.Run(); err != nil {
			return false, nil
		}
		return true, nil
	}
	return len(output) > 0, nil
}

// Install installs a tool using Scoop
func (w *windowsDeps) Install(toolName string) error {
	t := GetToolByName(toolName)
	if t == nil {
		return fmt.Errorf("unknown tool: %s", toolName)
	}

	parts := strings.Fields(t.WindowsCmd)
	if len(parts) < 2 {
		return fmt.Errorf("invalid install command for %s", toolName)
	}

	cmd := exec.Command(parts[0], parts[1:]...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to install %s: %s", toolName, string(output))
	}
	return nil
}

// CheckAndInstall verifies and installs missing dependencies
func (w *windowsDeps) CheckAndInstall(ctx context.Context, opts InstallOptions) (*InstallResult, error) {
	result := &InstallResult{
		Failed: make(map[string]error),
	}

	// Detect package manager
	pm, err := w.DetectPackageManager()
	if err != nil {
		result.Failed["package-manager"] = fmt.Errorf("no package manager available: %w", err)
		return result, nil
	}

	fmt.Printf("Using package manager: %s\n", pm)

	// Check each tool
	for _, tool := range RequiredTools {
		select {
		case <-ctx.Done():
			return result, ctx.Err()
		default:
		}

		installed, err := w.IsInstalled(tool.Name)
		if err != nil {
			result.Failed[tool.Name] = err
			continue
		}

		if installed {
			result.Skipped = append(result.Skipped, tool.Name)
			if opts.Verbose {
				fmt.Printf("  - %s: already installed\n", tool.Name)
			}
			continue
		}

		// Install the tool
		if opts.Verbose {
			fmt.Printf("  - %s: installing...\n", tool.Name)
		}

		if err := w.Install(tool.Name); err != nil {
			result.Failed[tool.Name] = err
			if opts.Verbose {
				fmt.Printf("  - %s: FAILED - %v\n", tool.Name, err)
			}
			continue
		}

		result.Installed = append(result.Installed, tool.Name)
		if opts.Verbose {
			fmt.Printf("  - %s: installed\n", tool.Name)
		}
	}

	return result, nil
}

// GetRequiredTools returns the list of required tools
func (w *windowsDeps) GetRequiredTools() []Tool {
	return RequiredTools
}

// DetectPlatform returns the current platform
func DetectPlatform() string {
	return "windows" // Could be extended to detect dynamically
}
