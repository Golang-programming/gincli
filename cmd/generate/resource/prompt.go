// cmd/generate/resource/prompt.go
package resource

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
)

func promptForValues() {
	if !skipPrompts {
		promptForResourceName()
		promptForEndpoints()
	}
}

func promptForResourceName() {
	if resourceName == "" {
		fmt.Printf("Enter resource name (e.g., [%s]): ", color.New(color.Faint).Sprint("user"))
		fmt.Scanln(&resourceName)
		if resourceName == "" {
			fmt.Println("Resource name is required.")
			ResourceCmd.Help()
			os.Exit(1)
		}
	}
}

func promptForEndpoints() {
	if !createEndpoints {
		prompt := &survey.Confirm{
			Message: "Do you want to create standard CRUD endpoints?",
			Default: false,
		}
		survey.AskOne(prompt, &createEndpoints)
	}
}
