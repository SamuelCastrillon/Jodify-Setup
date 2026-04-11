package dependencies

// Tool represents a system tool that can be installed
type Tool struct {
	Name        string
	Description string
	WindowsCmd  string // Scoop install command
	MacCmd      string // Brew install command
}

// List of required tools for Neovim plugins
var RequiredTools = []Tool{
	{
		Name:        "gcc",
		Description: "C compiler for Treesitter",
		WindowsCmd:  "scoop install gcc",
		MacCmd:      "brew install gcc",
	},
	{
		Name:        "ripgrep",
		Description: "Fast search for Telescope",
		WindowsCmd:  "scoop install ripgrep",
		MacCmd:      "brew install ripgrep",
	},
	{
		Name:        "fd",
		Description: "Fast alternative to find",
		WindowsCmd:  "scoop install fd",
		MacCmd:      "brew install fd",
	},
	{
		Name:        "fzf",
		Description: "Fuzzy finder",
		WindowsCmd:  "scoop install fzf",
		MacCmd:      "brew install fzf",
	},
	{
		Name:        "zoxide",
		Description: "Smart cd",
		WindowsCmd:  "scoop install zoxide",
		MacCmd:      "brew install zoxide",
	},
	{
		Name:        "lazygit",
		Description: "TUI for Git",
		WindowsCmd:  "scoop install lazygit",
		MacCmd:      "brew install lazygit",
	},
}

// GetToolByName returns a tool by its name
func GetToolByName(name string) *Tool {
	for i := range RequiredTools {
		if RequiredTools[i].Name == name {
			return &RequiredTools[i]
		}
	}
	return nil
}

// GetInstallCommand returns the appropriate install command for the platform
func (t *Tool) GetInstallCommand(platform string) string {
	switch platform {
	case "windows":
		return t.WindowsCmd
	case "darwin":
		return t.MacCmd
	default:
		return ""
	}
}
