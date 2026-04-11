package dependencies

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
)

// darwinDeps implements DependenciesManager for macOS using Homebrew
type darwinDeps struct{}

// NewDarwinDeps creates a new macOS dependencies manager
func NewDarwinDeps() DependenciesManager {
	return &darwinDeps{}
}

// DetectPackageManager returns the available package manager on macOS
func (d *darwinDeps) DetectPackageManager() (string, error) {
	cmd := exec.Command("brew", "--version")
	if err := cmd.Run(); err == nil {
		return "brew", nil
	}
	return "", fmt.Errorf("no package manager found (homebrew not available)")
}

// IsInstalled checks if a tool is installed on macOS
func (d *darwinDeps) IsInstalled(tool string) (bool, error) {
	cmd := exec.Command(tool, "--version")
	output, err := cmd.Output()
	if err != nil {
		cmd = exec.Command("brew", "list", tool)
		if err := cmd.Run(); err != nil {
			return false, nil
		}
		return true, nil
	}
	return len(output) > 0, nil
}

// Install installs a tool using Homebrew
func (d *darwinDeps) Install(toolName string) error {
	t := GetToolByName(toolName)
	if t == nil {
		return fmt.Errorf("unknown tool: %s", toolName)
	}

	parts := strings.Fields(t.MacCmd)
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

// CheckAndInstall verifies and installs missing dependencies on macOS
func (d *darwinDeps) CheckAndInstall(ctx context.Context, opts InstallOptions) (*InstallResult, error) {
	result := &InstallResult{
		Failed: make(map[string]error),
	}

	pm, err := d.DetectPackageManager()
	if err != nil {
		result.Failed["package-manager"] = fmt.Errorf("no package manager available: %w", err)
		return result, nil
	}

	fmt.Printf("Using package manager: %s\n", pm)

	for _, tool := range RequiredTools {
		select {
		case <-ctx.Done():
			return result, ctx.Err()
		default:
		}

		installed, err := d.IsInstalled(tool.Name)
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

		if opts.Verbose {
			fmt.Printf("  - %s: installing...\n", tool.Name)
		}

		if err := d.Install(tool.Name); err != nil {
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
func (d *darwinDeps) GetRequiredTools() []Tool {
	return RequiredTools
}
