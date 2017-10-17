package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd is the base of all comands
var RootCmd = &cobra.Command{
	Use:           "ptt",
	Short:         "Start server",
	Long:          `Start the ptt api service.`,
	SilenceErrors: true,
}
