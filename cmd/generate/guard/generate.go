// cmd/generate/guard/generate.go
package guard

import (
	"fmt"
	"path/filepath"

	"github.com/golang-programming/gincli/utils"
)

func generateGuardFile(projectDir string) {
	templatePath := "templates/others/guard.go.tpl"
	outputPath := filepath.Join(projectDir, fmt.Sprintf("%s.guard.go", guardName))
	config := map[string]string{
		"Module":              utils.DetectModuleName(),
		"GuardName":           utils.ConvertToSnakeCase(guardName),
		"CapitalizeGuardName": utils.Capitalize(guardName),
	}
	utils.GenerateFileFromTemplate(templatePath, outputPath, config)
}
