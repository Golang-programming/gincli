// cmd/generate/controller/generate.go
package controller

import (
	"fmt"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/golang-programming/gincli/utils"
	"github.com/spf13/cobra"
)

func createController(cmd *cobra.Command, args []string) {
	controllerName = args[0]

	promptForValues()

	projectDir := filepath.Join(".", "app", "controllers")
	utils.CreateDirectories([]string{projectDir})

	generateControllerFile(projectDir)

	fmt.Println(color.New(color.FgGreen).Sprint("Controller generated successfully"))
}

func generateControllerFile(projectDir string) {
	templatePath := "templates/generate/controller.go.tpl"
	outputPath := filepath.Join(projectDir, fmt.Sprintf("%s_controller.go", controllerName))
	config := map[string]string{
		"Module":                   utils.DetectModuleName(),
		"ControllerName":           controllerName,
		"CapitalizeControllerName": utils.Capitalize(controllerName),
	}
	utils.GenerateFileFromTemplate(templatePath, outputPath, config)
}
