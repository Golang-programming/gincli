// ./cmd/template/prompt.go
package template

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
)

func promptForValues() {
	promptForTemplate()
	promptForDBConfig()
}

func promptForTemplate() {
	if templateChoice == "" {
		prompt := &survey.Select{
			Message: fmt.Sprintf("Select a template you want to use: [%s]", color.New(color.Faint).Sprint(defaultTemplateChoice)),
			Options: availableTemplates,
		}
		survey.AskOne(prompt, &templateChoice)
	}
}

func promptForDBConfig() {
	if dbUsername == "" {
		prompt := &survey.Input{
			Message: fmt.Sprintf("Enter DB username [%s]: ", color.New(color.Faint).Sprint(defaultDBUsername)),
			Default: defaultDBUsername,
		}
		survey.AskOne(prompt, &dbUsername)
	}
	if dbPassword == "" {
		prompt := &survey.Password{
			Message: fmt.Sprintf("Enter DB password [%s]: ", color.New(color.Faint).Sprint(defaultDBPassword)),
		}
		survey.AskOne(prompt, &dbPassword)
	}
	if dbName == "" {
		prompt := &survey.Input{
			Message: fmt.Sprintf("Enter DB name [%s]: ", color.New(color.Faint).Sprint(defaultDBName)),
			Default: defaultDBName,
		}
		survey.AskOne(prompt, &dbName)
	}
	if dbHost == "" {
		prompt := &survey.Input{
			Message: fmt.Sprintf("Enter DB host [%s]: ", color.New(color.Faint).Sprint(defaultDBHost)),
			Default: defaultDBHost,
		}
		survey.AskOne(prompt, &dbHost)
	}
	if dbPort == "" {
		prompt := &survey.Input{
			Message: fmt.Sprintf("Enter DB port [%s]: ", color.New(color.Faint).Sprint(defaultMySQLPort)),
			Default: defaultMySQLPort,
		}
		survey.AskOne(prompt, &dbPort)
	}
}
