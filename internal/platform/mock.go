package platform

import (
	"os"
)

// MockPlatform is a mock implementation of Platform for testing
type MockPlatform struct {
	HomeDirFunc     func() (string, error)
	ConfigDirFunc   func() (string, error)
	DetectShellFunc func() (string, error)
	CheckPrereqFunc func() (*PrereqResult, error)
	OS              string
	HomeDir         string
	ConfigDir       string
	Shell           string
	NeovimInstalled bool
	NeovimVersion   string
	GitInstalled    bool
	GitVersion      string
	ErrToReturn     error
}

// NewMockPlatform creates a new MockPlatform with default values
func NewMockPlatform() *MockPlatform {
	return &MockPlatform{
		HomeDir:         "/tmp/home",
		ConfigDir:       "/tmp/home/.config/jodify",
		Shell:           "bash",
		NeovimInstalled: true,
		NeovimVersion:   "v0.10.0",
		GitInstalled:    true,
		GitVersion:      "git version 2.40.0",
		OS:              "darwin",
	}
}

// GetHomeDir returns the mocked home directory
func (m *MockPlatform) GetHomeDir() (string, error) {
	if m.HomeDirFunc != nil {
		return m.HomeDirFunc()
	}
	if m.ErrToReturn != nil {
		return "", m.ErrToReturn
	}
	return m.HomeDir, nil
}

// GetConfigDir returns the mocked config directory
func (m *MockPlatform) GetConfigDir() (string, error) {
	if m.ConfigDirFunc != nil {
		return m.ConfigDirFunc()
	}
	if m.ErrToReturn != nil {
		return "", m.ErrToReturn
	}
	return m.ConfigDir, nil
}

// DetectShell returns the mocked shell
func (m *MockPlatform) DetectShell() (string, error) {
	if m.DetectShellFunc != nil {
		return m.DetectShellFunc()
	}
	if m.ErrToReturn != nil {
		return "", m.ErrToReturn
	}
	return m.Shell, nil
}

// CheckPrerequisites returns the mocked prerequisite check result
func (m *MockPlatform) CheckPrerequisites() (*PrereqResult, error) {
	if m.CheckPrereqFunc != nil {
		return m.CheckPrereqFunc()
	}
	if m.ErrToReturn != nil {
		return nil, m.ErrToReturn
	}
	return &PrereqResult{
		NeovimInstalled: m.NeovimInstalled,
		NeovimVersion:   m.NeovimVersion,
		GitInstalled:    m.GitInstalled,
		GitVersion:      m.GitVersion,
		Errors:          []error{},
	}, nil
}

// GetOS returns the mocked OS
func (m *MockPlatform) GetOS() string {
	return m.OS
}

// Ensure MockPlatform implements Platform
var _ Platform = (*MockPlatform)(nil)

// MockFileSystem is a mock for filesystem operations
type MockFileSystem struct {
	Files       map[string][]byte
	Directories map[string]bool
	ErrToReturn error
}

// NewMockFileSystem creates a new MockFileSystem
func NewMockFileSystem() *MockFileSystem {
	return &MockFileSystem{
		Files:       make(map[string][]byte),
		Directories: make(map[string]bool),
	}
}

// WriteFile writes data to a file in the mock filesystem
func (m *MockFileSystem) WriteFile(path string, data []byte, perm os.FileMode) error {
	if m.ErrToReturn != nil {
		return m.ErrToReturn
	}
	m.Files[path] = data
	return nil
}

// ReadFile reads data from a file in the mock filesystem
func (m *MockFileSystem) ReadFile(path string) ([]byte, error) {
	if m.ErrToReturn != nil {
		return nil, m.ErrToReturn
	}
	data, ok := m.Files[path]
	if !ok {
		return nil, os.ErrNotExist
	}
	return data, nil
}

// MkdirAll creates a directory in the mock filesystem
func (m *MockFileSystem) MkdirAll(path string, perm os.FileMode) error {
	if m.ErrToReturn != nil {
		return m.ErrToReturn
	}
	m.Directories[path] = true
	return nil
}

// Exists checks if a file or directory exists in the mock filesystem
func (m *MockFileSystem) Exists(path string) bool {
	_, fileExists := m.Files[path]
	_, dirExists := m.Directories[path]
	return fileExists || dirExists
}
