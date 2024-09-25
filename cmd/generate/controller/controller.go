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
)

var ControllerCmd = &cobra.Command{
	Use:   "controller <name>",
	Short: "Generate a new controller",
	Run:   createController,
}

func createController(cmd *cobra.Command, args []string) {
	controllerName = args[0]

	promptForValues()

	projectDir := filepath.Join(".", "app", "modules", controllerName, "controllers")

	generateControllerFile(projectDir)

	fmt.Println(color.New(color.FgGreen).Sprint("Controller generated successfully"))
}
