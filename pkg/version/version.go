package version

import "fmt"

// Version is the current version of jodify-setup
// Set via ldflags at build time (e.g., -X github.com/SamuelCastrillon/Jodify-Setup/pkg/version.Version=v1.0.0)
var Version = "v0.0.0"

// Commit is the git commit hash
// Set via ldflags at build time (e.g., -X github.com/SamuelCastrillon/Jodify-Setup/pkg/version.Commit=$(git rev-parse --short HEAD))
var Commit = "dev"

// Date is the build timestamp
// Set via ldflags at build time (e.g., -X github.com/SamuelCastrillon/Jodify-Setup/pkg/version.Date=$(date -u +"%Y-%m-%dT%H:%M:%SZ"))
var Date = "now"

// String returns the version information as a formatted string
func String() string {
	return fmt.Sprintf("jodify-setup %s (commit: %s, date: %s)", Version, Commit, Date)
}
