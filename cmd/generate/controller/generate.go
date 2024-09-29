// ./cmd/generate/controller/generate.go
package controller

import (
	"fmt"
	"path/filepath"

	"github.com/golang-programming/gincli/utils"
)

func generateControllerFile(projectDir string) {
	// Template path relative to the embedded FS
	templatePath := "templates/others/controller.go.tpl"
	outputPath := filepath.Join(projectDir, fmt.Sprintf("%s.controller.go", utils.ConvertToSnakeCase(controllerName)))
	config := map[string]string{
		"Module":                   utils.DetectModuleName(),
		"ControllerName":           utils.ConvertToSnakeCase(controllerName),
		"CapitalizeControllerName": utils.ToPascalCase(controllerName),
	}
	utils.GenerateFileFromTemplate(templatePath, outputPath, config)
}
