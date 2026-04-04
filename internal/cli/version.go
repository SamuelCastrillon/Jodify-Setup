package cli

import (
	"fmt"

	"github.com/SamuelCastrillon/Jodify-Setup/pkg/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Long:  "Display the version, commit hash, and build date of jodify-setup",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.String())
	},
}
