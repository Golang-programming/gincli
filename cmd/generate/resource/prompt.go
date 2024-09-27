// ./cmd/generate/resource/prompt.go
package resource

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
)

func promptForValues() {
	promptForTransport()
	promptForEndpoints()
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

func promptForEndpoints() {
	if !createEndpoint {
		prompt := &survey.Confirm{
			Message: "Do you want to create endpoints?",
			Default: true,
		}
		survey.AskOne(prompt, &createEndpoint)
	}
}
