package main

import (
	"fmt"
	"os"

	"github.com/SamuelCastrillon/Jodify-Setup/internal/cli"
	"github.com/SamuelCastrillon/Jodify-Setup/pkg/version"
)

func main() {
	// Check for --version flag first
	if len(os.Args) > 1 && (os.Args[1] == "--version" || os.Args[1] == "-v") {
		fmt.Println(version.String())
		os.Exit(0)
	}

	// Run CLI
	if err := cli.Execute(); err != nil {
		if err == cli.ErrCheckFlag {
			// --check was handled, exit code already set
			return
		}
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
