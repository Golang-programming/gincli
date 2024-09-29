// ./cmd/generate/controller/controller.go
package controller

import (
	"fmt"
	"path/filepath"

	"github.com/golang-programming/gincli/utils"
	"github.com/spf13/cobra"
)

var (
	controllerName string
	controllerPath string
)

var ControllerCmd = &cobra.Command{
	Use:     "controller <name> [path]",
	Short:   "Generate a new controller",
	Aliases: []string{"c"},
	Args:    cobra.MinimumNArgs(1), // Ensures at least one argument (controller name) is passed
	Run:     createController,
}

func init() {
	// Add shorthand flags if needed in the future
}

func createController(cmd *cobra.Command, args []string) {
	controllerName = utils.ConvertToSnakeCase(args[0])
	defaultProjectDir := filepath.Join(".", "app", "modules", controllerName, "controllers")

	// Check if a custom path is provided as the second argument
	if len(args) > 1 {
		controllerPath = args[1]
	} else {
		controllerPath = defaultProjectDir
	}

	generateControllerFile(controllerPath)

	utils.LogSuccess(fmt.Sprintf("Controller generated successfully at path: %s", controllerPath))
}
