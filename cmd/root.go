// ./cmd/root.go
package cmd

import (
	"os"

	"github.com/golang-programming/gincli/cmd/generate"
	"github.com/golang-programming/gincli/cmd/new"
	"github.com/golang-programming/gincli/cmd/template"
	"github.com/golang-programming/gincli/utils"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// CustomHelpFunc defines a custom help function that displays commands in a table.
func CustomHelpFunc(cmd *cobra.Command, args []string) {
	if cmd.HasParent() {
		// If the command has a parent, use the default help
		cmd.Parent().HelpFunc()(cmd, args)
		return
	}

	utils.LogInfo("Gin CLI - Scaffold Your Gin Application\n")
	utils.LogInfo("Usage:")
	utils.LogInfo("  gin [command] [flags]\n")
	utils.LogInfo("Available Commands:")

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Command", "Aliases", "Description"})

	commands := []struct {
		Name        string
		Aliases     string
		Description string
	}{
		{"new", "n, create", "Create a new Gin application with a project structure"},
		{"generate", "g", "Generate resources, controllers, guards, routes, and services"},
		{"template", "t", "Load application startup templates"},
	}

	for _, cmd := range commands {
		table.Append([]string{cmd.Name, cmd.Aliases, cmd.Description})
	}

	table.Render()

	utils.LogInfo("\nUse \"gin [command] --help\" for more information about a command.")
}

var rootCmd = &cobra.Command{
	Use:   "gin",
	Short: "CLI to generate Gin applications with different components",
	Long:  `A CLI tool that helps generate Gin applications and its components.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	// Silence Cobra's default error and usage messages
	rootCmd.SilenceErrors = true
	rootCmd.SilenceUsage = true

	if err := rootCmd.Execute(); err != nil {
		// Log the error in red color
		utils.LogError(err.Error())
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(template.TemplateCmd)
	rootCmd.AddCommand(new.NewCmd)
	rootCmd.AddCommand(generate.GenerateCmd) // Add the generate command

	// Set the custom help function
	rootCmd.SetHelpFunc(CustomHelpFunc)

	// Handle unknown commands by setting a custom function
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true}) // Hide the default help command
	rootCmd.SetUsageTemplate("")                         // Remove default usage template

	// Add a global flag for help if needed
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Display help information")
}
