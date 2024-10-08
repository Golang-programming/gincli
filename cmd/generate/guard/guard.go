// ./cmd/generate/guard/guard.go
package guard

import (
	"path/filepath"

	"github.com/golang-programming/gincli/utils"
	"github.com/spf13/cobra"
)

var (
	guardName string
	guardPath string
)

var GuardCmd = &cobra.Command{
	Use:     "guard <name> [path]",
	Short:   "Generate a new guard",
	Aliases: []string{"gd"},
	Args:    cobra.MinimumNArgs(1),
	Run:     createGuard,
}

func init() {
	// Add shorthand flags if needed in the future
}

func createGuard(cmd *cobra.Command, args []string) {
	guardName = utils.ConvertToSnakeCase(args[0])
	defaultPath := filepath.Join(".", "app", "guards", guardName)

	if len(args) > 1 {
		guardPath = args[1]
	} else {
		guardPath = defaultPath
	}

	generateGuardFile(guardPath)

	utils.LogSuccess("Guard generated successfully")
}
