// cmd/generate/controller/prompt.go
package controller

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func promptForValues() {
	if !skipPrompts {
		promptForControllerName()
	}
}

func promptForControllerName() {
	if controllerName == "" {
		fmt.Printf("Enter controller name (e.g., [%s]): ", color.New(color.Faint).Sprint("user"))
		fmt.Scanln(&controllerName)
		if controllerName == "" {
			fmt.Println("Controller name is required.")
			ControllerCmd.Help()
			os.Exit(1)
		}
	}
}
