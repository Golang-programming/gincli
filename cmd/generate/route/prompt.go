// cmd/generate/route/prompt.go
package route

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func promptForValues() {
	if !skipPrompts {
		promptForRouteName()
	}
}

func promptForRouteName() {
	if routeName == "" {
		fmt.Printf("Enter route name (e.g., [%s]): ", color.New(color.Faint).Sprint("user"))
		fmt.Scanln(&routeName)
		if routeName == "" {
			fmt.Println("Route name is required.")
			RouteCmd.Help()
			os.Exit(1)
		}
	}
}
