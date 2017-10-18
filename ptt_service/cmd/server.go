package cmd

import (
	"fmt"
	"ptt/http"

	"github.com/spf13/cobra"
)

var (
	// Port to run the server on
	Port = "9000"
)

func init() {
	RootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringVarP(&Port, "port", "p", "9000", "Server port to run on")
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the ptt server",
	Long:  `Starts the ptt server on the provided port or 9000 if not specified.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Starting ptt on port %s\n", Port)
		http.Server(Port)
	},
}
