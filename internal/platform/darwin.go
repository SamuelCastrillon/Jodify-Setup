package platform

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// darwin implements Platform for macOS
type darwin struct{}

// NewDarwin creates a new Darwin/macOS platform
func NewDarwin() Platform {
	return &darwin{}
}

// GetHomeDir returns the user's home directory on macOS
func (d *darwin) GetHomeDir() (string, error) {
	home := os.Getenv("HOME")
	if home == "" {
		return "", ErrHomeDirNotFound
	}
	return home, nil
}

// GetConfigDir returns the Neovim config directory on macOS
// Uses ~/.config with NVIM_APPNAME for isolation
func (d *darwin) GetConfigDir() (string, error) {
	// Support NVIM_APPNAME environment variable
	appName := os.Getenv("NVIM_APPNAME")
	if appName == "" {
		appName = "nvim"
	}

	home, err := d.GetHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, ".config", appName), nil
}

// DetectShell identifies the user's preferred shell on macOS
func (d *darwin) DetectShell() (string, error) {
	// Check for user's default shell
	shell := os.Getenv("SHELL")
	if shell != "" {
		parts := strings.Split(shell, "/")
		if len(parts) > 0 {
			return parts[len(parts)-1], nil
		}
	}

	// Default to zsh on modern macOS
	return "zsh", nil
}

// CheckPrerequisites checks if Neovim and Git are installed on macOS
func (d *darwin) CheckPrerequisites() (*PrereqResult, error) {
	result := &PrereqResult{
		Errors: []error{},
	}

	// Check Neovim
	nvimPath, err := findDarwinCommand("nvim")
	if err == nil {
		result.NeovimInstalled = true
		version, err := getDarwinCommandVersion(nvimPath, "--version")
		if err == nil {
			result.NeovimVersion = extractVersionLine(version)
		}
	}

	// Check Git
	gitPath, err := findDarwinCommand("git")
	if err == nil {
		result.GitInstalled = true
		version, err := getDarwinCommandVersion(gitPath, "--version")
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
func (d *darwin) GetOS() string {
	return "darwin"
}

// findDarwinCommand finds a command in PATH on macOS
func findDarwinCommand(name string) (string, error) {
	cmd := exec.Command("which", name)
	output, err := cmd.Output()
	if err != nil {
		return "", ErrCommandNotFound
	}

	path := strings.TrimSpace(string(output))
	if path == "" {
		return "", ErrCommandNotFound
	}

	return path, nil
}

// getDarwinCommandVersion gets the version output of a command
func getDarwinCommandVersion(path, arg string) (string, error) {
	cmd := exec.Command(path, arg)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}
