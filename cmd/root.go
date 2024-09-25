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
	if cmd.Name() == "gincli" {
		utils.LogInfo("Gin CLI - Scaffold Your Gin Application\n")
		utils.LogInfo("Usage:")
		utils.LogInfo("  gincli [command] [flags]\n")
		utils.LogInfo("Available Commands:")

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Command", "Description"})

		commands := []struct {
			Name        string
			Description string
		}{
			{"new", "Create a new Gin application with a project structure"},
			{"generate", "Generate resources, controllers, services, and routes"},
			{"template", "Load application startup template"},
		}

		for _, cmd := range commands {
			table.Append([]string{cmd.Name, cmd.Description})
		}

		table.Render()

		utils.LogInfo("\nUse \"gincli [command] --help\" for more information about a command.")
	} else {
		cmd.Parent().HelpFunc()(cmd, args)
	}
}

var rootCmd = &cobra.Command{
	Use:   "gincli",
	Short: "CLI to generate Gin applications with different components",
	Long:  `A CLI tool that helps generate Gin applications and its components.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		utils.LogError(err.Error())
	}
}

func init() {
	rootCmd.AddCommand(template.TemplateCmd)
	rootCmd.AddCommand(new.NewCmd)
	rootCmd.AddCommand(generate.GenerateCmd) // Add the generate command
	// Set the custom help function
	rootCmd.SetHelpFunc(CustomHelpFunc)
}
