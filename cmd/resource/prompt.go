package resource

import (
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2"

	"github.com/fatih/color"
)

func promptForValues() {
	promptForResourceName()
	promptForTransport()
}

func promptForResourceName() {
	if resourceName == "" {
		fmt.Printf("Enter resource name (e.g, [%s]): ", color.New(color.Faint).Sprint("user"))
		fmt.Scanln(&resourceName)
		if resourceName == "" {
			log.Fatalf("Please enter a resource name")
		}
	}
}

func promptForTransport() {
	if transport == "" {
		prompt := &survey.Select{
			Message: fmt.Sprintf("What transport layer do you use? [%s]", color.New(color.Faint).Sprint(defaultTransport)),
			Options: availableTransports,
		}
		survey.AskOne(prompt, &transport)
	}
}
