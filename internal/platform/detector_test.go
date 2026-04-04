package platform

import (
	"testing"
)

func TestDetectGoOS(t *testing.T) {
	tests := []struct {
		name     string
		goos     string
		expected string
		wantErr  bool
	}{
		{
			name:     "detects windows",
			goos:     "windows",
			expected: "windows",
			wantErr:  false,
		},
		{
			name:     "detects darwin",
			goos:     "darwin",
			expected: "darwin",
			wantErr:  false,
		},
		{
			name:     "rejects linux",
			goos:     "linux",
			expected: "",
			wantErr:  true,
		},
		{
			name:     "rejects freebsd",
			goos:     "freebsd",
			expected: "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DetectGoOS(tt.goos)
			if (err != nil) != tt.wantErr {
				t.Errorf("DetectGoOS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got.GetOS() != tt.expected {
				t.Errorf("DetectGoOS() = %v, want %v", got.GetOS(), tt.expected)
			}
		})
	}
}

func TestDetect(t *testing.T) {
	// Just verify it doesn't panic - it will detect based on runtime.GOOS
	p, err := Detect()
	if err == ErrUnsupportedPlatform {
		t.Logf("Current platform not supported: %v", err)
		return
	}
	if err != nil {
		t.Errorf("Detect() error = %v", err)
		return
	}
	if p == nil {
		t.Error("Detect() returned nil platform")
	}
}

func TestMockPlatformImplementsPlatform(t *testing.T) {
	mock := NewMockPlatform()
	var _ Platform = mock
}

func TestMockPlatformGetHomeDir(t *testing.T) {
	mock := NewMockPlatform()

	home, err := mock.GetHomeDir()
	if err != nil {
		t.Errorf("GetHomeDir() error = %v", err)
	}
	if home != "/tmp/home" {
		t.Errorf("GetHomeDir() = %v, want /tmp/home", home)
	}
}

func TestMockPlatformGetConfigDir(t *testing.T) {
	mock := NewMockPlatform()

	configDir, err := mock.GetConfigDir()
	if err != nil {
		t.Errorf("GetConfigDir() error = %v", err)
	}
	if configDir != "/tmp/home/.config/jodify" {
		t.Errorf("GetConfigDir() = %v, want /tmp/home/.config/jodify", configDir)
	}
}

func TestMockPlatformDetectShell(t *testing.T) {
	mock := NewMockPlatform()

	shell, err := mock.DetectShell()
	if err != nil {
		t.Errorf("DetectShell() error = %v", err)
	}
	if shell != "bash" {
		t.Errorf("DetectShell() = %v, want bash", shell)
	}
}

func TestMockPlatformCheckPrerequisites(t *testing.T) {
	mock := NewMockPlatform()

	result, err := mock.CheckPrerequisites()
	if err != nil {
		t.Errorf("CheckPrerequisites() error = %v", err)
	}
	if !result.NeovimInstalled {
		t.Error("CheckPrerequisites() NeovimInstalled = false, want true")
	}
	if !result.GitInstalled {
		t.Error("CheckPrerequisites() GitInstalled = false, want true")
	}
}

func TestMockPlatformGetOS(t *testing.T) {
	mock := NewMockPlatform()

	os := mock.GetOS()
	if os != "darwin" {
		t.Errorf("GetOS() = %v, want darwin", os)
	}
}

func TestMockPlatformWithCustomFuncs(t *testing.T) {
	mock := &MockPlatform{
		HomeDirFunc: func() (string, error) {
			return "/custom/home", nil
		},
		ConfigDirFunc: func() (string, error) {
			return "/custom/config", nil
		},
		DetectShellFunc: func() (string, error) {
			return "zsh", nil
		},
		CheckPrereqFunc: func() (*PrereqResult, error) {
			return &PrereqResult{
				NeovimInstalled: false,
				GitInstalled:    true,
			}, nil
		},
		OS: "windows",
	}

	home, err := mock.GetHomeDir()
	if err != nil {
		t.Errorf("GetHomeDir() error = %v", err)
	}
	if home != "/custom/home" {
		t.Errorf("GetHomeDir() = %v, want /custom/home", home)
	}

	configDir, err := mock.GetConfigDir()
	if err != nil {
		t.Errorf("GetConfigDir() error = %v", err)
	}
	if configDir != "/custom/config" {
		t.Errorf("GetConfigDir() = %v, want /custom/config", configDir)
	}

	shell, err := mock.DetectShell()
	if err != nil {
		t.Errorf("DetectShell() error = %v", err)
	}
	if shell != "zsh" {
		t.Errorf("DetectShell() = %v, want zsh", shell)
	}

	result, err := mock.CheckPrerequisites()
	if err != nil {
		t.Errorf("CheckPrerequisites() error = %v", err)
	}
	if result.NeovimInstalled {
		t.Error("CheckPrerequisites() NeovimInstalled = true, want false")
	}
	if !result.GitInstalled {
		t.Error("CheckPrerequisites() GitInstalled = false, want true")
	}

	if mock.GetOS() != "windows" {
		t.Errorf("GetOS() = %v, want windows", mock.GetOS())
	}
}

func TestMockPlatformErrorPropagation(t *testing.T) {
	mock := &MockPlatform{
		ErrToReturn: ErrHomeDirNotFound,
	}

	_, err := mock.GetHomeDir()
	if err != ErrHomeDirNotFound {
		t.Errorf("GetHomeDir() error = %v, want ErrHomeDirNotFound", err)
	}
}

func TestMockPlatformErrorPropagationOnConfigDir(t *testing.T) {
	mock := &MockPlatform{
		ErrToReturn: ErrHomeDirNotFound,
	}

	_, err := mock.GetConfigDir()
	if err != ErrHomeDirNotFound {
		t.Errorf("GetConfigDir() error = %v", err)
	}
}

func TestMockPlatformErrorPropagationOnShell(t *testing.T) {
	mock := &MockPlatform{
		ErrToReturn: ErrCommandNotFound,
	}

	_, err := mock.DetectShell()
	if err != ErrCommandNotFound {
		t.Errorf("DetectShell() error = %v", err)
	}
}

func TestMockPlatformPrereqResult(t *testing.T) {
	result := &PrereqResult{
		NeovimInstalled: true,
		NeovimVersion:   "v0.10.0",
		GitInstalled:    true,
		GitVersion:      "git version 2.40.0",
		Errors:          []error{},
	}

	if !result.NeovimInstalled {
		t.Error("NeovimInstalled should be true")
	}
	if result.NeovimVersion != "v0.10.0" {
		t.Errorf("NeovimVersion = %v, want v0.10.0", result.NeovimVersion)
	}
	if !result.GitInstalled {
		t.Error("GitInstalled should be true")
	}
	if result.GitVersion != "git version 2.40.0" {
		t.Errorf("GitVersion = %v, want git version 2.40.0", result.GitVersion)
	}
	if len(result.Errors) != 0 {
		t.Errorf("Errors should be empty, got %d", len(result.Errors))
	}
}

func TestMockPlatformPrereqResultWithErrors(t *testing.T) {
	result := &PrereqResult{
		NeovimInstalled: false,
		GitInstalled:    false,
		Errors:          []error{ErrNeovimNotFound, ErrGitNotFound},
	}

	if len(result.Errors) != 2 {
		t.Errorf("Expected 2 errors, got %d", len(result.Errors))
	}
}

func TestMockFileSystem(t *testing.T) {
	mfs := NewMockFileSystem()

	// Test WriteFile and ReadFile
	err := mfs.WriteFile("/test/file.txt", []byte("hello world"), 0644)
	if err != nil {
		t.Errorf("WriteFile() error = %v", err)
	}

	data, err := mfs.ReadFile("/test/file.txt")
	if err != nil {
		t.Errorf("ReadFile() error = %v", err)
	}
	if string(data) != "hello world" {
		t.Errorf("ReadFile() = %v, want 'hello world'", string(data))
	}

	// Test Exists
	if !mfs.Exists("/test/file.txt") {
		t.Error("Exists() should return true for existing file")
	}
	if mfs.Exists("/nonexistent") {
		t.Error("Exists() should return false for non-existent path")
	}

	// Test MkdirAll
	err = mfs.MkdirAll("/new/dir", 0755)
	if err != nil {
		t.Errorf("MkdirAll() error = %v", err)
	}
	if !mfs.Exists("/new/dir") {
		t.Error("MkdirAll() should create directory")
	}
}

func TestMockFileSystemError(t *testing.T) {
	mfs := &MockFileSystem{
		ErrToReturn: ErrCommandNotFound,
	}

	err := mfs.WriteFile("/test/file.txt", []byte("test"), 0644)
	if err != ErrCommandNotFound {
		t.Errorf("WriteFile() error = %v, want ErrCommandNotFound", err)
	}

	_, err = mfs.ReadFile("/test/file.txt")
	if err != ErrCommandNotFound {
		t.Errorf("ReadFile() error = %v, want ErrCommandNotFound", err)
	}

	err = mfs.MkdirAll("/new/dir", 0755)
	if err != ErrCommandNotFound {
		t.Errorf("MkdirAll() error = %v, want ErrCommandNotFound", err)
	}
}

func TestNewDarwin(t *testing.T) {
	d := NewDarwin()
	if d.GetOS() != "darwin" {
		t.Errorf("GetOS() = %v, want darwin", d.GetOS())
	}
}

func TestNewWindows(t *testing.T) {
	w := NewWindows()
	if w.GetOS() != "windows" {
		t.Errorf("GetOS() = %v, want windows", w.GetOS())
	}
}
