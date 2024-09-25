// cmd/generate/route/route.go
package route

import (
	"github.com/spf13/cobra"
)

var (
	routeName   string
	skipPrompts bool
)

var RouteCmd = &cobra.Command{
	Use:     "route <name>",
	Short:   "Generate a new route",
	Args:    cobra.ExactArgs(1),
	Run:     createRoute,
	Aliases: []string{"rt"},
}

func init() {
	RouteCmd.Flags().BoolVarP(&skipPrompts, "yes", "y", false, "Skip prompts and use default values")
}
