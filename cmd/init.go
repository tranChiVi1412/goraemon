package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [service-name]",
	Short: "Initialize a new Go microservice",
	Long:  `This command initializes a new Go microservice with the specified name. It creates a directory structure and necessary files for the service.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		svcName := args[0]
		fmt.Printf("Initializing new Go microservice: %s\n", svcName)
		// TODO: Call template engine to render the service template
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
