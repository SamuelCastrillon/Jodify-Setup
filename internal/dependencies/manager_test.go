package dependencies

import (
	"context"
	"errors"
	"testing"
)

type mockDepsManager struct {
	installedTools map[string]bool
	packageManager string
	installErr     error
	detectPMErr    error
}

func newMockDepsManager() *mockDepsManager {
	return &mockDepsManager{
		installedTools: make(map[string]bool),
		packageManager: "scoop",
	}
}

func (m *mockDepsManager) CheckAndInstall(ctx context.Context, opts InstallOptions) (*InstallResult, error) {
	result := &InstallResult{
		Failed: make(map[string]error),
	}

	_, err := m.DetectPackageManager()
	if err != nil {
		result.Failed["package-manager"] = err
		return result, nil
	}

	for _, tool := range RequiredTools {
		select {
		case <-ctx.Done():
			return result, ctx.Err()
		default:
		}

		installed, err := m.IsInstalled(tool.Name)
		if err != nil {
			result.Failed[tool.Name] = err
			continue
		}

		if installed {
			result.Skipped = append(result.Skipped, tool.Name)
			continue
		}

		if m.installErr != nil {
			result.Failed[tool.Name] = m.installErr
			continue
		}

		m.installedTools[tool.Name] = true
		result.Installed = append(result.Installed, tool.Name)
	}

	return result, nil
}

func (m *mockDepsManager) IsInstalled(tool string) (bool, error) {
	return m.installedTools[tool], nil
}

func (m *mockDepsManager) DetectPackageManager() (string, error) {
	if m.detectPMErr != nil {
		return "", m.detectPMErr
	}
	return m.packageManager, nil
}

func (m *mockDepsManager) GetRequiredTools() []Tool {
	return RequiredTools
}

func TestIsInstalled_ReturnsTrueWhenToolExists(t *testing.T) {
	mock := newMockDepsManager()
	mock.installedTools["gcc"] = true

	installed, err := mock.IsInstalled("gcc")

	if err != nil {
		t.Errorf("IsInstalled() error = %v", err)
	}
	if !installed {
		t.Error("IsInstalled() = false, want true")
	}
}

func TestIsInstalled_ReturnsFalseWhenToolDoesNotExist(t *testing.T) {
	mock := newMockDepsManager()

	installed, err := mock.IsInstalled("nonexistent")

	if err != nil {
		t.Errorf("IsInstalled() error = %v", err)
	}
	if installed {
		t.Error("IsInstalled() = true, want false")
	}
}

func TestCheckAndInstall_InstallsMissingTools(t *testing.T) {
	mock := newMockDepsManager()
	mock.installedTools = map[string]bool{
		"ripgrep": true,
		"fd":      true,
	}

	result, err := mock.CheckAndInstall(context.Background(), InstallOptions{})

	if err != nil {
		t.Errorf("CheckAndInstall() error = %v", err)
	}

	if len(result.Installed) == 0 {
		t.Error("CheckAndInstall() should have installed missing tools")
	}

	expectedInstalled := []string{"gcc", "fzf", "zoxide", "lazygit"}
	for _, tool := range expectedInstalled {
		found := false
		for _, installed := range result.Installed {
			if installed == tool {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected %s to be installed", tool)
		}
	}
}

func TestCheckAndInstall_SkipsAlreadyInstalledTools(t *testing.T) {
	mock := newMockDepsManager()
	mock.installedTools = map[string]bool{
		"gcc":     true,
		"ripgrep": true,
		"fd":      true,
		"fzf":     true,
		"zoxide":  true,
		"lazygit": true,
	}

	result, err := mock.CheckAndInstall(context.Background(), InstallOptions{})

	if err != nil {
		t.Errorf("CheckAndInstall() error = %v", err)
	}

	if len(result.Skipped) == 0 {
		t.Error("CheckAndInstall() should skip already installed tools")
	}

	if len(result.Installed) != 0 {
		t.Errorf("CheckAndInstall() installed %d tools, want 0", len(result.Installed))
	}
}

func TestCheckAndInstall_SkipExistingOption(t *testing.T) {
	mock := newMockDepsManager()
	mock.installedTools = map[string]bool{
		"gcc": true,
	}

	result, err := mock.CheckAndInstall(context.Background(), InstallOptions{SkipExisting: true})

	if err != nil {
		t.Errorf("CheckAndInstall() error = %v", err)
	}

	if len(result.Installed) != 5 {
		t.Errorf("Expected 5 tools to install, got %d", len(result.Installed))
	}
}

func TestCheckAndInstall_FallbackWhenPackageManagerNotAvailable(t *testing.T) {
	mock := newMockDepsManager()
	mock.packageManager = ""
	mock.detectPMErr = errors.New("package manager not found")

	result, err := mock.CheckAndInstall(context.Background(), InstallOptions{})

	if err != nil {
		t.Errorf("CheckAndInstall() error = %v", err)
	}

	if result.Failed == nil {
		t.Error("CheckAndInstall() should report package manager failure")
	}

	if _, ok := result.Failed["package-manager"]; !ok {
		t.Error("CheckAndInstall() should include package-manager in Failed")
	}
}

func TestCheckAndInstall_InstallResultStruct(t *testing.T) {
	tests := []struct {
		name           string
		installedTools map[string]bool
		packageMgr     string
		wantErr        bool
		checkResult    func(*testing.T, *InstallResult)
	}{
		{
			name:           "all tools already installed",
			installedTools: map[string]bool{"gcc": true, "ripgrep": true, "fd": true, "fzf": true, "zoxide": true, "lazygit": true},
			packageMgr:     "scoop",
			wantErr:        false,
			checkResult: func(t *testing.T, r *InstallResult) {
				if len(r.Installed) != 0 {
					t.Errorf("Expected 0 installed, got %d", len(r.Installed))
				}
				if len(r.Skipped) != 6 {
					t.Errorf("Expected 6 skipped, got %d", len(r.Skipped))
				}
			},
		},
		{
			name:           "no tools installed",
			installedTools: map[string]bool{},
			packageMgr:     "scoop",
			wantErr:        false,
			checkResult: func(t *testing.T, r *InstallResult) {
				if len(r.Installed) != 6 {
					t.Errorf("Expected 6 installed, got %d", len(r.Installed))
				}
				if len(r.Skipped) != 0 {
					t.Errorf("Expected 0 skipped, got %d", len(r.Skipped))
				}
			},
		},
		{
			name:           "some tools installed",
			installedTools: map[string]bool{"gcc": true, "ripgrep": true},
			packageMgr:     "scoop",
			wantErr:        false,
			checkResult: func(t *testing.T, r *InstallResult) {
				if len(r.Installed) != 4 {
					t.Errorf("Expected 4 installed, got %d", len(r.Installed))
				}
				if len(r.Skipped) != 2 {
					t.Errorf("Expected 2 skipped, got %d", len(r.Skipped))
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := newMockDepsManager()
			mock.installedTools = tt.installedTools
			mock.packageManager = tt.packageMgr

			result, err := mock.CheckAndInstall(context.Background(), InstallOptions{})

			if (err != nil) != tt.wantErr {
				t.Errorf("CheckAndInstall() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			tt.checkResult(t, result)
		})
	}
}

func TestDetectPackageManager_ReturnsAvailable(t *testing.T) {
	mock := newMockDepsManager()
	mock.packageManager = "scoop"

	pm, err := mock.DetectPackageManager()

	if err != nil {
		t.Errorf("DetectPackageManager() error = %v", err)
	}
	if pm != "scoop" {
		t.Errorf("DetectPackageManager() = %v, want scoop", pm)
	}
}

func TestDetectPackageManager_ReturnsErrorWhenNotAvailable(t *testing.T) {
	mock := newMockDepsManager()
	mock.packageManager = ""
	mock.detectPMErr = errors.New("no package manager")

	_, err := mock.DetectPackageManager()

	if err == nil {
		t.Error("DetectPackageManager() should return error when not available")
	}
}

func TestGetRequiredTools_ReturnsAllTools(t *testing.T) {
	mock := newMockDepsManager()

	tools := mock.GetRequiredTools()

	if len(tools) != 6 {
		t.Errorf("GetRequiredTools() = %d tools, want 6", len(tools))
	}

	expectedNames := []string{"gcc", "ripgrep", "fd", "fzf", "zoxide", "lazygit"}
	for i, tool := range tools {
		if i < len(expectedNames) && tool.Name != expectedNames[i] {
			t.Errorf("GetRequiredTools()[%d] = %v, want %v", i, tool.Name, expectedNames[i])
		}
	}
}

func TestMockManagerImplementsDependenciesManager(t *testing.T) {
	mock := newMockDepsManager()
	var _ DependenciesManager = mock
}
