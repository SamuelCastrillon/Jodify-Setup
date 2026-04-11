package cli

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/SamuelCastrillon/Jodify-Setup/internal/platform"
	"github.com/spf13/cobra"
)

var devOpenCmd = &cobra.Command{
	Use:   "open",
	Short: "Open Neovim with Jodify config",
	Long: `Open Neovim with the Jodify configuration.
	
Uses NVIM_APPNAME=jodify to isolate the config from your existing Neovim setup.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		p, err := platform.Detect()
		if err != nil {
			return fmt.Errorf("failed to detect platform: %w", err)
		}

		// Check if Neovim is installed
		result, err := p.CheckPrerequisites()
		if err != nil {
			return fmt.Errorf("failed to check prerequisites: %w", err)
		}

		if !result.NeovimInstalled {
			return fmt.Errorf("Neovim is not installed. Please install Neovim first")
		}

		// Get config directory
		configDir, err := p.GetConfigDir()
		if err != nil {
			return fmt.Errorf("failed to get config directory: %w", err)
		}

		// Check if config exists
		if _, err := os.Stat(configDir); os.IsNotExist(err) {
			return fmt.Errorf("no config found at %s. Run 'jodify-setup sync' first", configDir)
		}

		// Find Neovim executable
		nvimPath, err := findNeovim()
		if err != nil {
			return fmt.Errorf("failed to find Neovim: %w", err)
		}

		fmt.Printf("Opening Neovim with config at: %s\n", configDir)
		fmt.Println("Tip: This uses NVIM_APPNAME=jodify for isolation")

		// Set NVIM_APPNAME and run Neovim
		cmdEnv := os.Environ()
		cmdEnv = append(cmdEnv, "NVIM_APPNAME=jodify")

		nvimCmd := exec.Command(nvimPath)
		nvimCmd.Env = cmdEnv
		nvimCmd.Stdin = os.Stdin
		nvimCmd.Stdout = os.Stdout
		nvimCmd.Stderr = os.Stderr

		return nvimCmd.Run()
	},
}

// findNeovim tries to find the Neovim executable
func findNeovim() (string, error) {
	// Try common locations
	paths := []string{
		"nvim",
		"nvim.exe",
	}

	for _, p := range paths {
		if _, err := exec.LookPath(p); err == nil {
			return p, nil
		}
	}

	// Check common Windows install locations
	commonPaths := []string{
		"C:\\Program Files\\Neovim\\bin\\nvim.exe",
		"C:\\Program Files (x86)\\Neovim\\bin\\nvim.exe",
	}

	for _, p := range commonPaths {
		if _, err := os.Stat(p); err == nil {
			return p, nil
		}
	}

	return "", fmt.Errorf("Neovim not found. Please install Neovim")
}

func init() {
	devOpenCmd.Flags().String("profile", "jodify", "Neovim profile to use")
}
