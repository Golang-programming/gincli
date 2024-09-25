// cmd/generate/route/prompt.go
package route

import (
	"fmt"

	"github.com/fatih/color"
)

func promptForValues() {
	if routeName == "" {
		fmt.Printf("Enter route name (e.g., [%s]): ", color.New(color.Faint).Sprint("user"))
		fmt.Scanln(&routeName)
		if routeName == "" {
			fmt.Println("Route name is required.")
		}
	}
}
