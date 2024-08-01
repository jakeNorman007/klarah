package cmd

import (
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command {
	Use:   "klarah",
	Short: "Scaffolding CLI tool that creates a quick backend for a multitude of projects",
	Long: "Klarah is a template CLI tool that scaffolds a backend using Golang and is ready to go with minimal effort",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
