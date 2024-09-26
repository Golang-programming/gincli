// ./cmd/generate/generate.go
package generate

import (
	"os"

	"github.com/golang-programming/gincli/cmd/generate/controller"
	"github.com/golang-programming/gincli/cmd/generate/guard"
	"github.com/golang-programming/gincli/cmd/generate/resource"
	"github.com/golang-programming/gincli/cmd/generate/route"
	"github.com/golang-programming/gincli/cmd/generate/service"
	"github.com/golang-programming/gincli/utils"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// GenerateCmd is the parent command for all generate subcommands
var GenerateCmd = &cobra.Command{
	Use:     "generate",
	Short:   "Generate resources, controllers, guards, routes, and services",
	Long:    `Generate scaffolding for resources, controllers, guards, routes, and services similar to NestJS CLI.`,
	Aliases: []string{"g"},
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// CustomGenerateHelpFunc defines a custom help function for the generate command.
func CustomGenerateHelpFunc(cmd *cobra.Command, args []string) {
	utils.LogInfo("Generate Command - Generate resources, controllers, guards, routes, and services.\n")
	utils.LogInfo("Usage:")
	utils.LogInfo("  gincli generate [subcommand] [flags]\n")
	utils.LogInfo("Available Subcommands:")

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Subcommand", "Aliases", "Description"})

	subcommands := []struct {
		Name        string
		Aliases     string
		Description string
	}{
		{"controller", "c", "Generate a new controller"},
		{"guard", "gd", "Generate a new guard"},
		{"resource", "r", "Generate a new resource"},
		{"route", "rt", "Generate a new route"},
		{"service", "s", "Generate a new service"},
	}

	for _, sc := range subcommands {
		table.Append([]string{sc.Name, sc.Aliases, sc.Description})
	}

	table.Render()

	utils.LogInfo("\nUse \"gincli generate [subcommand] --help\" for more information about a subcommand.")
}

func init() {
	GenerateCmd.AddCommand(resource.ResourceCmd)
	GenerateCmd.AddCommand(guard.GuardCmd)
	GenerateCmd.AddCommand(controller.ControllerCmd)
	GenerateCmd.AddCommand(service.ServiceCmd)
	GenerateCmd.AddCommand(route.RouteCmd)
	// Set the custom help function
	GenerateCmd.SetHelpFunc(CustomGenerateHelpFunc)
}
