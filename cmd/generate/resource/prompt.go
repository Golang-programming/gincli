package resource

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"

	"github.com/fatih/color"
)

func promptForValues() {
	if transport == "" {
		prompt := &survey.Select{
			Message: fmt.Sprintf("What transport layer do you use? [%s]", color.New(color.Faint).Sprint(defaultTransport)),
			Options: availableTransports,
		}
		survey.AskOne(prompt, &transport)
	}
}
