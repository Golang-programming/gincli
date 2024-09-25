// cmd/generate/route/route.go
package route

import (
	"fmt"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	routeName string
)

var RouteCmd = &cobra.Command{
	Use:   "route <name>",
	Short: "Generate a new route",
	Run:   createRoute,
}

func createRoute(cmd *cobra.Command, args []string) {
	routeName = args[0]

	promptForValues()

	projectDir := filepath.Join(".", "app", "modules", routeName)

	generateRouteFile(projectDir)

	fmt.Println(color.New(color.FgGreen).Sprint("Route generated successfully"))
}
