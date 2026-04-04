package platform

import "errors"

// Platform defines the interface for platform-specific operations
type Platform interface {
	// GetHomeDir returns the user's home directory
	GetHomeDir() (string, error)

	// GetConfigDir returns the Neovim config directory path
	// Uses NVIM_APPNAME for isolation: ~/.config/jodify on Unix
	GetConfigDir() (string, error)

	// DetectShell identifies the user's preferred shell
	DetectShell() (string, error)

	// CheckPrerequisites verifies Neovim and Git are installed
	CheckPrerequisites() (*PrereqResult, error)

	// GetOS returns the operating system name
	GetOS() string
}

// PrereqResult holds the result of checking system prerequisites
type PrereqResult struct {
	NeovimInstalled bool
	NeovimVersion   string
	GitInstalled    bool
	GitVersion      string
	Errors          []error
}

// ErrUnsupportedPlatform is returned when the platform is not supported
var ErrUnsupportedPlatform = errors.New("unsupported platform: linux is not supported in v1")

// ErrNeovimNotFound is returned when Neovim is not installed
var ErrNeovimNotFound = errors.New("neovim not found")

// ErrGitNotFound is returned when Git is not installed
var ErrGitNotFound = errors.New("git not found")

// ErrHomeDirNotFound is returned when the home directory cannot be determined
var ErrHomeDirNotFound = errors.New("home directory not found")

// ErrCommandNotFound is returned when a command is not found in PATH
var ErrCommandNotFound = errors.New("command not found")
