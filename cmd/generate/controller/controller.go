// cmd/generate/controller/controller.go
package controller

import (
	"fmt"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	controllerName string
	controllerPath string
)

var ControllerCmd = &cobra.Command{
	Use:   "controller <name> [path]",
	Short: "Generate a new controller",
	Args:  cobra.MinimumNArgs(1), // Ensures at least one argument (controller name) is passed
	Run:   createController,
}

func createController(cmd *cobra.Command, args []string) {
	controllerName = args[0]
	defaultProjectDir := filepath.Join(".", "app", "modules", controllerName, "controllers")

	// Check if a custom path is provided as the second argument
	if len(args) > 1 {
		controllerPath = args[1]
	} else {
		controllerPath = defaultProjectDir
	}

	generateControllerFile(controllerPath)

	fmt.Println(color.New(color.FgGreen).Sprint("Controller generated successfully at path: " + controllerPath))
}
