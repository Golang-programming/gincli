// cmd/generate/route/generate.go
package route

import (
	"fmt"
	"path/filepath"

	"github.com/golang-programming/gincli/utils"
)

func generateRouteFile(projectDir string) {
	templatePath := "templates/others/route.go.tpl"
	outputPath := filepath.Join(projectDir, fmt.Sprintf("%s.route.go", routeName))

	config := map[string]string{
		"Module":              utils.DetectModuleName(),
		"RouteName":           utils.ConvertToSnakeCase(routeName),
		"CapitalizeRouteName": utils.Capitalize(routeName),
	}
	utils.GenerateFileFromTemplate(templatePath, outputPath, config)
}