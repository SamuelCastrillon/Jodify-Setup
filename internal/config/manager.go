package config

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const (
	// DefaultTimeout is the default HTTP timeout
	DefaultTimeout = 30 * time.Second
	// MaxRetries is the maximum number of retries for download
	MaxRetries = 3
	// DefaultReleaseURL is the default config release URL
	DefaultReleaseURL = "https://github.com/jodify/jodify-config/releases/latest/download/jodify-config.zip"
)

// Manager handles configuration download and extraction
type Manager interface {
	// Download fetches the config zip from the given URL
	Download(ctx context.Context, url string) (string, error)

	// Extract unpacks the config zip to the target directory
	Extract(zipPath, targetDir string) error

	// VerifyChecksum validates the downloaded file against expected hash
	VerifyChecksum(filePath, expectedHash string) error

	// GetLatestReleaseURL returns the download URL for the latest release
	GetLatestReleaseURL() string
}

// manager implements Manager
type manager struct {
	httpClient *http.Client
	cacheDir   string
}

// NewManager creates a new config Manager
func NewManager() Manager {
	return &manager{
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
	}
}

// Download fetches the config zip from the given URL with retry logic
func (m *manager) Download(ctx context.Context, url string) (string, error) {
	var lastErr error

	for attempt := 1; attempt <= MaxRetries; attempt++ {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		default:
		}

		// Create temp file
		tmpFile, err := os.CreateTemp("", "jodify-config-*.zip")
		if err != nil {
			return "", fmt.Errorf("failed to create temp file: %w", err)
		}
		tmpPath := tmpFile.Name()
		tmpFile.Close()
		defer os.Remove(tmpPath)

		// Download with timeout
		downloadCtx, cancel := context.WithTimeout(ctx, DefaultTimeout)
		defer cancel()

		req, err := http.NewRequestWithContext(downloadCtx, http.MethodGet, url, nil)
		if err != nil {
			lastErr = fmt.Errorf("failed to create request: %w", err)
			continue
		}

		resp, err := m.httpClient.Do(req)
		if err != nil {
			lastErr = fmt.Errorf("download failed: %w", err)
			time.Sleep(time.Second * time.Duration(attempt))
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			lastErr = fmt.Errorf("download failed with status: %d", resp.StatusCode)
			time.Sleep(time.Second * time.Duration(attempt))
			continue
		}

		// Copy response body to file
		file, err := os.OpenFile(tmpPath, os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			return "", fmt.Errorf("failed to open temp file: %w", err)
		}

		_, err = io.Copy(file, resp.Body)
		file.Close()
		if err != nil {
			lastErr = fmt.Errorf("failed to write to temp file: %w", err)
			continue
		}

		return tmpPath, nil
	}

	return "", fmt.Errorf("download failed after %d attempts: %w", MaxRetries, lastErr)
}

// Extract unpacks the config zip to the target directory
func (m *manager) Extract(zipPath, targetDir string) error {
	// Open zip file
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return fmt.Errorf("failed to open zip file: %w", err)
	}
	defer r.Close()

	// Create target directory
	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return fmt.Errorf("failed to create target directory: %w", err)
	}

	// Extract all files
	for _, f := range r.File {
		path := filepath.Join(targetDir, f.Name)

		// Security: prevent zip slip vulnerability
		if !strings.HasPrefix(path, filepath.Clean(targetDir)+string(os.PathSeparator)) {
			return fmt.Errorf("invalid zip entry: %s", f.Name)
		}

		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(path, 0755); err != nil {
				return fmt.Errorf("failed to create directory: %w", err)
			}
			continue
		}

		// Create parent directory
		parent := filepath.Dir(path)
		if err := os.MkdirAll(parent, 0755); err != nil {
			return fmt.Errorf("failed to create parent directory: %w", err)
		}

		// Extract file
		outFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.FileInfo().Mode())
		if err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}
		defer outFile.Close()

		inFile, err := f.Open()
		if err != nil {
			return fmt.Errorf("failed to open zip entry: %w", err)
		}
		defer inFile.Close()

		if _, err := io.Copy(outFile, inFile); err != nil {
			return fmt.Errorf("failed to extract file: %w", err)
		}
	}

	return nil
}

// VerifyChecksum validates the downloaded file against expected hash
func (m *manager) VerifyChecksum(filePath, expectedHash string) error {
	// TODO: Implement SHA256 checksum verification
	_ = filePath
	_ = expectedHash
	return nil
}

// GetLatestReleaseURL returns the download URL for the latest release
func (m *manager) GetLatestReleaseURL() string {
	return DefaultReleaseURL
}
