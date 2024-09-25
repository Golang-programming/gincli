// cmd/generate/service/prompt.go
package service

import (
	"fmt"

	"github.com/fatih/color"
)

func promptForValues() {
	if !skipPrompts {
		promptForServiceName()
	}
}

func promptForServiceName() {
	if serviceName == "" {
		fmt.Printf("Enter service name (e.g., [%s]): ", color.New(color.Faint).Sprint("user"))
		fmt.Scanln(&serviceName)
		if serviceName == "" {
			fmt.Println("Service name is required.")
		}
	}
}
