// ./cmd/generate/service/generate.go
package service

import (
	"fmt"
	"path/filepath"

	"github.com/golang-programming/gincli/utils"
)

func generateServiceFile(projectDir string) {
	templatePath := "templates/others/service.go.tpl"
	outputPath := filepath.Join(projectDir, fmt.Sprintf("%s.service.go", utils.ConvertToSnakeCase(serviceName)))
	config := map[string]string{
		"Module":                utils.DetectModuleName(),
		"ServiceName":           utils.ConvertToSnakeCase(serviceName),
		"CapitalizeServiceName": utils.Capitalize(serviceName),
	}
	utils.GenerateFileFromTemplate(templatePath, outputPath, config)
}
