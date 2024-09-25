// cmd/generate/controller/prompt.go
package controller

import (
	"fmt"

	"github.com/fatih/color"
)

func promptForValues() {
	if controllerName == "" {
		fmt.Printf("Enter controller name (e.g., [%s]): ", color.New(color.Faint).Sprint("user"))
		fmt.Scanln(&controllerName)
		if controllerName == "" {
			fmt.Println("Service name is required.")
		}
	}
}
