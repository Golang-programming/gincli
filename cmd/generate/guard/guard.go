// cmd/generate/guard/guard.go
package guard

import (
	"fmt"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	guardName string
	guardPath string
)

var GuardCmd = &cobra.Command{
	Use:   "guard <name> [path]",
	Short: "Generate a new guard",
	Args:  cobra.MinimumNArgs(1),
	Run:   createGuard,
}

func createGuard(cmd *cobra.Command, args []string) {
	guardName = args[0]
	defaultPath := filepath.Join(".", "app", guardName)

	if len(args) > 1 {
		guardPath = args[1]
	} else {
		guardPath = defaultPath
	}

	generateGuardFile(guardPath)

	fmt.Println(color.New(color.FgGreen).Sprint("Guard generated successfully"))
}
