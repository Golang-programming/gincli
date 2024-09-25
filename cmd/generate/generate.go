// cmd/generate/generate.go
package generate

import (
	"github.com/golang-programming/gincli/cmd/generate/controller"
	"github.com/golang-programming/gincli/cmd/generate/route"
	"github.com/golang-programming/gincli/cmd/generate/service"
	"github.com/spf13/cobra"
)

var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate resources, controllers, services, and routes",
	Long:  `Generate scaffolding for resources, controllers, services, and routes similar to NestJS CLI.`,
}

func init() {
	// GenerateCmd.AddCommand(.ResourceCmd)
	GenerateCmd.AddCommand(controller.ControllerCmd)
	GenerateCmd.AddCommand(service.ServiceCmd)
	GenerateCmd.AddCommand(route.RouteCmd)
}
