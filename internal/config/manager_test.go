package config

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/SamuelCastrillon/Jodify-Setup/pkg/version"
)

// TestDownloadSuccess tests successful download
func TestDownloadSuccess(t *testing.T) {
	// Create test server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test content"))
	}))
	defer server.Close()

	mgr := NewManager()

	path, err := mgr.Download(context.Background(), server.URL)
	if err != nil {
		t.Errorf("Download() error = %v", err)
	}
	defer os.Remove(path)

	if path == "" {
		t.Error("Download() returned empty path")
	}
}

// TestDownloadRetry tests retry logic on failure
func TestDownloadRetry(t *testing.T) {
	attempts := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		attempts++
		if attempts < 3 {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("test content"))
	}))
	defer server.Close()

	mgr := NewManager()

	path, err := mgr.Download(context.Background(), server.URL)
	if err != nil {
		t.Errorf("Download() error = %v", err)
	}
	defer os.Remove(path)

	if attempts != 3 {
		t.Errorf("expected 3 attempts, got %d", attempts)
	}
}

// TestDownloadAllFail tests download when all retries fail
func TestDownloadAllFail(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	defer server.Close()

	mgr := NewManager()

	_, err := mgr.Download(context.Background(), server.URL)
	if err == nil {
		t.Error("Download() expected error when all retries fail")
	}
}

// TestGetLatestReleaseURL tests getting release URL
func TestGetLatestReleaseURL(t *testing.T) {
	mgr := NewManager()

	url := mgr.GetLatestReleaseURL()
	// In dev mode (version=v0.0.0), should return latest release URL
	if url != "https://github.com/SamuelCastrillon/Jodify-Setup/releases/latest/download/jodify-config.zip" {
		t.Errorf("GetLatestReleaseURL() = %v, want latest release URL", url)
	}
}

// TestGetLatestReleaseURLWithVersion tests URL generation for release builds
func TestGetLatestReleaseURLWithVersion(t *testing.T) {
	// Save original version
	origVersion := version.Version
	defer func() { version.Version = origVersion }()

	// Set a release version
	version.Version = "v0.1.0"

	mgr := NewManager()
	url := mgr.GetLatestReleaseURL()

	expected := "https://github.com/SamuelCastrillon/Jodify-Setup/releases/download/v0.1.0/jodify-config-0.1.0.zip"
	if url != expected {
		t.Errorf("GetLatestReleaseURL() = %v\nwant %v", url, expected)
	}
}

// TestVerifyChecksum tests checksum verification
func TestVerifyChecksum(t *testing.T) {
	mgr := NewManager()

	// Should not error (placeholder implementation)
	err := mgr.VerifyChecksum("test.txt", "abc123")
	if err != nil {
		t.Errorf("VerifyChecksum() error = %v", err)
	}
}
