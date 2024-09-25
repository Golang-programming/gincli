// ./cmd/generate/route/route.go
package route

import (
	"fmt"
	"path/filepath"

	"github.com/golang-programming/gincli/utils"
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

func init() {
	// Add shorthand flags if needed in the future
}

func createRoute(cmd *cobra.Command, args []string) {
	routeName = args[0]
	defaultProjectDir := filepath.Join(".", "app", "modules", routeName, "routes")

	// Check if a custom path is provided as the second argument
	if len(args) > 1 {
		routePath = args[1]
	} else {
		routePath = defaultProjectDir
	}

	generateRouteFile(routePath)

	utils.LogSuccess(fmt.Sprintf("Route generated successfully at path: %s", routePath))
}
