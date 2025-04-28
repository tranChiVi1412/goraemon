package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goraemon",
	Short: "Goraemon - A assistant for your Go building scafford microservices",
	Long:  `Goraemon is a command line tool that helps you to create and manage Go microservices easily.`,
}

func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error executing command:", err)
		os.Exit(1)
	}
	return nil
}
