// cmd/generate/service/generate.go
package service

import (
	"fmt"
	"path/filepath"

	"github.com/golang-programming/gincli/utils"
)

func generateServiceFile(projectDir string) {
	templatePath := "templates/generate/service.go.tpl"
	outputPath := filepath.Join(projectDir, fmt.Sprintf("%s_service.go", serviceName))
	config := map[string]string{
		"Module":                utils.DetectModuleName(),
		"ServiceName":           serviceName,
		"CapitalizeServiceName": utils.Capitalize(serviceName),
	}
	utils.GenerateFileFromTemplate(templatePath, outputPath, config)
}
