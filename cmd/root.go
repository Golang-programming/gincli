package cmd

import (
    "os"

    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "gincli",
    Short: "A CLI tool to generate Gin applications",
    Long:  `gincli is a CLI tool to create and manage Gin web applications.`,
}

// Execute executes the root command.
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}

func init() {
    // Here you can define flags and configuration settings.
}
