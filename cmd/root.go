// cmd/root.go
package cmd

import (
	"os"

	"github.com/golang-programming/gincli/cmd/generate"
	"github.com/golang-programming/gincli/cmd/new"
	"github.com/golang-programming/gincli/cmd/template"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gincli",
	Short: "CLI to generate Gin applications with different components",
	Long:  `A CLI tool that helps generate Gin applications and its components.`,
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
	rootCmd.AddCommand(generate.GenerateCmd) // Add the generate command
	// Removed: rootCmd.AddCommand(helpCmd)
}
