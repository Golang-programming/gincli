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

var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate resources, controllers, services, and routes",
	Long:  `Generate scaffolding for resources, controllers, services, and routes similar to NestJS CLI.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// CustomGenerateHelpFunc defines a custom help function for the generate command.
func CustomGenerateHelpFunc(cmd *cobra.Command, args []string) {
	utils.LogInfo("Generate Command - Generate resources, controllers, services, and routes.\n")
	utils.LogInfo("Usage:")
	utils.LogInfo("  gincli generate [subcommand] [flags]\n")
	utils.LogInfo("Available Subcommands:")

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Subcommand", "Description"})

	subcommands := []struct {
		Name        string
		Description string
	}{
		{"controller", "Generate a new controller"},
		{"guard", "Generate a new guard"},
		{"resource", "Generate a new resource"},
		{"route", "Generate a new route"},
		{"service", "Generate a new service"},
	}

	for _, sc := range subcommands {
		table.Append([]string{sc.Name, sc.Description})
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
