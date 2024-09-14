package cmd

import (
	"os"

	"github.com/golang-programming/gincli/cmd/new"
	"github.com/golang-programming/gincli/cmd/template"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gincli",
	Short: "CLI to generate Gin web applications with different components.",
	Long:  `A CLI tool for building scalable Gin applications, with the ability to generate controllers, services, models, middleware, routes, and more.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(template.TemplateCmd)
	rootCmd.AddCommand(new.NewCmd)
	// Removed: rootCmd.AddCommand(helpCmd)
}
