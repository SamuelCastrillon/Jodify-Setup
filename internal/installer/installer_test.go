package installer

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/SamuelCastrillon/Jodify-Setup/internal/platform"
)

// mockConfigManager implements a minimal config manager for testing
type mockConfigManager struct {
	downloadURL  string
	downloadErr  error
	extractErr   error
	downloadPath string
}

func (m *mockConfigManager) Download(ctx context.Context, url string) (string, error) {
	if m.downloadErr != nil {
		return "", m.downloadErr
	}
	if m.downloadPath != "" {
		return m.downloadPath, nil
	}
	tmpFile, err := os.CreateTemp("", "test-config-*.zip")
	if err != nil {
		return "", err
	}
	tmpFile.Close()
	return tmpFile.Name(), nil
}

func (m *mockConfigManager) Extract(zipPath, targetDir string) error {
	return m.extractErr
}

func (m *mockConfigManager) VerifyChecksum(filePath, expectedHash string) error {
	return nil
}

func (m *mockConfigManager) GetLatestReleaseURL() string {
	return m.downloadURL
}

// TestInstallerInstall tests the Install method
func TestInstallerInstall(t *testing.T) {
	tmpHome := t.TempDir()
	configDir := filepath.Join(tmpHome, ".config", "jodify")

	mockPlatform := &platform.MockPlatform{
		HomeDir:         tmpHome,
		ConfigDir:       configDir,
		NeovimInstalled: true,
		NeovimVersion:   "v0.10.0",
		GitInstalled:    true,
		GitVersion:      "git version 2.40.0",
	}

	mockConfigMgr := &mockConfigManager{}
	inst := NewWithConfig(mockPlatform, mockConfigMgr)

	opts := InstallOptions{
		Force:      true,
		SkipBackup: true,
	}

	err := inst.Install(context.Background(), opts)
	if err != nil {
		t.Errorf("Install() error = %v", err)
	}
}

// TestInstallerInstallNeovimMissing tests install with Neovim missing
func TestInstallerInstallNeovimMissing(t *testing.T) {
	tmpHome := t.TempDir()
	configDir := filepath.Join(tmpHome, ".config", "jodify")

	mockPlatform := &platform.MockPlatform{
		HomeDir:         tmpHome,
		ConfigDir:       configDir,
		NeovimInstalled: false,
		GitInstalled:    true,
	}

	mockConfigMgr := &mockConfigManager{}
	inst := NewWithConfig(mockPlatform, mockConfigMgr)

	opts := InstallOptions{}

	err := inst.Install(context.Background(), opts)
	if err == nil {
		t.Error("Install() expected error when Neovim is missing")
	}
}

// TestInstallerInstallGitMissing tests install with Git missing
func TestInstallerInstallGitMissing(t *testing.T) {
	tmpHome := t.TempDir()
	configDir := filepath.Join(tmpHome, ".config", "jodify")

	mockPlatform := &platform.MockPlatform{
		HomeDir:         tmpHome,
		ConfigDir:       configDir,
		NeovimInstalled: true,
		GitInstalled:    false,
	}

	mockConfigMgr := &mockConfigManager{}
	inst := NewWithConfig(mockPlatform, mockConfigMgr)

	opts := InstallOptions{}

	err := inst.Install(context.Background(), opts)
	if err == nil {
		t.Error("Install() expected error when Git is missing")
	}
}

// TestInstallerUpdate tests the Update method
func TestInstallerUpdate(t *testing.T) {
	tmpHome := t.TempDir()
	configDir := filepath.Join(tmpHome, ".config", "jodify")

	os.MkdirAll(configDir, 0755)

	mockPlatform := &platform.MockPlatform{
		HomeDir:         tmpHome,
		ConfigDir:       configDir,
		NeovimInstalled: true,
		GitInstalled:    true,
	}

	mockConfigMgr := &mockConfigManager{}
	inst := NewWithConfig(mockPlatform, mockConfigMgr)

	opts := UpdateOptions{}

	err := inst.Update(context.Background(), opts)
	if err != nil {
		t.Errorf("Update() error = %v", err)
	}
}

// TestInstallerUpdateNoInstall tests update without prior installation
func TestInstallerUpdateNoInstall(t *testing.T) {
	tmpHome := t.TempDir()
	configDir := filepath.Join(tmpHome, ".config", "jodify")

	mockPlatform := &platform.MockPlatform{
		HomeDir:         tmpHome,
		ConfigDir:       configDir,
		NeovimInstalled: true,
		GitInstalled:    true,
	}

	mockConfigMgr := &mockConfigManager{}
	inst := NewWithConfig(mockPlatform, mockConfigMgr)

	opts := UpdateOptions{}

	err := inst.Update(context.Background(), opts)
	if err == nil {
		t.Error("Update() expected error when no prior installation")
	}
}

// TestInstallerUninstall tests the Uninstall method
func TestInstallerUninstall(t *testing.T) {
	tmpHome := t.TempDir()
	configDir := filepath.Join(tmpHome, ".config", "jodify")

	os.MkdirAll(configDir, 0755)

	mockPlatform := &platform.MockPlatform{
		HomeDir:         tmpHome,
		ConfigDir:       configDir,
		NeovimInstalled: true,
		GitInstalled:    true,
	}

	mockConfigMgr := &mockConfigManager{}
	inst := NewWithConfig(mockPlatform, mockConfigMgr)

	opts := UninstallOptions{}

	err := inst.Uninstall(context.Background(), opts)
	if err != nil {
		t.Errorf("Uninstall() error = %v", err)
	}

	if _, err := os.Stat(configDir); !os.IsNotExist(err) {
		t.Error("Uninstall() should remove config directory")
	}
}

// TestInstallerUninstallNoConfig tests uninstall with no config
func TestInstallerUninstallNoConfig(t *testing.T) {
	tmpHome := t.TempDir()
	configDir := filepath.Join(tmpHome, ".config", "jodify")

	mockPlatform := &platform.MockPlatform{
		HomeDir:         tmpHome,
		ConfigDir:       configDir,
		NeovimInstalled: true,
		GitInstalled:    true,
	}

	mockConfigMgr := &mockConfigManager{}
	inst := NewWithConfig(mockPlatform, mockConfigMgr)

	opts := UninstallOptions{}

	err := inst.Uninstall(context.Background(), opts)
	if err != nil {
		t.Errorf("Uninstall() error = %v", err)
	}
}

// TestInstallerUninstallWithBackup tests uninstall with backup restore
func TestInstallerUninstallWithBackup(t *testing.T) {
	tmpHome := t.TempDir()
	configDir := filepath.Join(tmpHome, ".config", "jodify")
	backupDir := configDir + ".backup"

	os.MkdirAll(configDir, 0755)
	os.MkdirAll(backupDir, 0755)

	mockPlatform := &platform.MockPlatform{
		HomeDir:         tmpHome,
		ConfigDir:       configDir,
		NeovimInstalled: true,
		GitInstalled:    true,
	}

	mockConfigMgr := &mockConfigManager{}
	inst := NewWithConfig(mockPlatform, mockConfigMgr)

	opts := UninstallOptions{
		RestoreBackup: true,
	}

	err := inst.Uninstall(context.Background(), opts)
	if err != nil {
		t.Errorf("Uninstall() error = %v", err)
	}
}

// TestConfigExists tests config existence check
func TestConfigExists(t *testing.T) {
	tmpDir := t.TempDir()
	configDir := filepath.Join(tmpDir, "config")

	exists, err := ConfigExists(configDir)
	if err != nil {
		t.Errorf("ConfigExists() error = %v", err)
	}
	if exists {
		t.Error("ConfigExists() = true, want false for non-existent config")
	}

	os.MkdirAll(configDir, 0755)

	exists, err = ConfigExists(configDir)
	if err != nil {
		t.Errorf("ConfigExists() error = %v", err)
	}
	if !exists {
		t.Error("ConfigExists() = false, want true for existing config")
	}
}

// TestGetConfigBackupDir tests backup directory naming
func TestGetConfigBackupDir(t *testing.T) {
	configDir := "/home/user/.config/jodify"
	backupDir := GetConfigBackupDir(configDir)

	expected := configDir + ".backup"
	if backupDir != expected {
		t.Errorf("GetConfigBackupDir() = %v, want %v", backupDir, expected)
	}
}
