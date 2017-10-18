package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
	"ptt/domain"
)

func init() {
	RootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version",
	Long:  `Display version and build information about ptt.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("ptt %s\n", domain.Version)
		fmt.Printf("  Build date: %s\n", domain.BuildDate)
		fmt.Printf("  Built with: %s\n", runtime.Version())
	},
}
