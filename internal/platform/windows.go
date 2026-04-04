package platform

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// windows implements Platform for Windows
type windows struct{}

// NewWindows creates a new Windows platform
func NewWindows() Platform {
	return &windows{}
}

// GetHomeDir returns the user's home directory on Windows
func (w *windows) GetHomeDir() (string, error) {
	home := os.Getenv("USERPROFILE")
	if home == "" {
		return "", ErrHomeDirNotFound
	}
	return home, nil
}

// GetConfigDir returns the Neovim config directory on Windows
// Uses %LOCALAPPDATA% for better Windows integration
func (w *windows) GetConfigDir() (string, error) {
	// Support NVIM_APPNAME environment variable
	appName := os.Getenv("NVIM_APPNAME")
	if appName == "" {
		appName = "nvim"
	}

	localAppData := os.Getenv("LOCALAPPDATA")
	if localAppData == "" {
		// Fallback to USERPROFILE
		home, err := w.GetHomeDir()
		if err != nil {
			return "", err
		}
		return filepath.Join(home, "AppData", "Local", appName), nil
	}

	return filepath.Join(localAppData, appName), nil
}

// DetectShell identifies the user's preferred shell on Windows
func (w *windows) DetectShell() (string, error) {
	// Check for PowerShell first (preferred on modern Windows)
	pshome := os.Getenv("PSModulePath")
	if pshome != "" {
		return "powershell", nil
	}

	// Fallback to cmd
	return "cmd", nil
}

// CheckPrerequisites checks if Neovim and Git are installed on Windows
func (w *windows) CheckPrerequisites() (*PrereqResult, error) {
	result := &PrereqResult{
		Errors: []error{},
	}

	// Check Neovim
	nvimPath, err := findWindowsCommand("nvim")
	if err == nil {
		result.NeovimInstalled = true
		version, err := getWindowsCommandVersion(nvimPath, "--version")
		if err == nil {
			result.NeovimVersion = extractVersionLine(version)
		}
	}

	// Check Git
	gitPath, err := findWindowsCommand("git")
	if err == nil {
		result.GitInstalled = true
		version, err := getWindowsCommandVersion(gitPath, "--version")
		if err == nil {
			result.GitVersion = extractVersionLine(version)
		}
	}

	if !result.NeovimInstalled {
		result.Errors = append(result.Errors, ErrNeovimNotFound)
	}
	if !result.GitInstalled {
		result.Errors = append(result.Errors, ErrGitNotFound)
	}

	return result, nil
}

// GetOS returns the operating system name
func (w *windows) GetOS() string {
	return "windows"
}

// findWindowsCommand finds a command in PATH on Windows
func findWindowsCommand(name string) (string, error) {
	cmd := exec.Command("where", name)
	output, err := cmd.Output()
	if err != nil {
		return "", ErrCommandNotFound
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	if len(lines) == 0 {
		return "", ErrCommandNotFound
	}

	return strings.TrimSpace(lines[0]), nil
}

// getWindowsCommandVersion gets the version output of a command
func getWindowsCommandVersion(path, arg string) (string, error) {
	cmd := exec.Command(path, arg)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// extractVersionLine extracts the version string from command output
func extractVersionLine(output string) string {
	lines := strings.Split(output, "\n")
	if len(lines) > 0 {
		return strings.TrimSpace(lines[0])
	}
	return "unknown"
}
