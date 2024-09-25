// cmd/generate/route/generate.go
package route

import (
	"fmt"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/golang-programming/gincli/utils"
	"github.com/spf13/cobra"
)

func createRoute(cmd *cobra.Command, args []string) {
	routeName = args[0]

	promptForValues()

	projectDir := filepath.Join(".", "app", "routes")
	utils.CreateDirectories([]string{projectDir})

	generateRouteFile(projectDir)

	fmt.Println(color.New(color.FgGreen).Sprint("Route generated successfully"))
}

func generateRouteFile(projectDir string) {
	templatePath := "templates/generate/route.go.tpl"
	outputPath := filepath.Join(projectDir, fmt.Sprintf("%s_route.go", routeName))
	config := map[string]string{
		"Module":              utils.DetectModuleName(),
		"RouteName":           routeName,
		"CapitalizeRouteName": utils.Capitalize(routeName),
	}
	utils.GenerateFileFromTemplate(templatePath, outputPath, config)
}
