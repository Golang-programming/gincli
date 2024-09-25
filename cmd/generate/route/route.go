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
	routePath string
)

var RouteCmd = &cobra.Command{
	Use:   "route <name> [path]",
	Short: "Generate a new route",
	Args:  cobra.MinimumNArgs(1),
	Run:   createRoute,
}

func createRoute(cmd *cobra.Command, args []string) {
	routeName = args[0]
	defaultProjectDir := filepath.Join(".", "app", "modules", routeName)

	// Check if a custom path is provided as the second argument
	if len(args) > 1 {
		routePath = args[1]
	} else {
		routePath = defaultProjectDir
	}

	generateRouteFile(routePath)

	fmt.Println(color.New(color.FgGreen).Sprint("Route generated successfully at path: " + routePath))
}
