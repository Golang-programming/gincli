// ./cmd/new/prompt.go
package new

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
)

func promptForValues() {
	promptForDBType()
	if strings.ToLower(dbType) == "mongodb" {
		promptForMongoDBUri()
	} else {
		if strings.ToLower(dbType) != "sqlite" {
			promptForDBConfig()
		}
	}
}

func promptForMongoDBUri() {
	prompt := &survey.Input{
		Message: "MongoDB URI: ",
	}
	survey.AskOne(prompt, &mongodbUri)
}

func promptForDBType() {
	if dbType == "" {
		prompt := &survey.Select{
			Message: fmt.Sprintf("Select your database: [%s]", color.New(color.Faint).Sprint(defaultDBType)),
			Options: availableDBTypes,
		}
		survey.AskOne(prompt, &dbType)
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
		defaultPort := defaultMySQLPort
		if strings.ToLower(dbType) == "postgres" {
			defaultPort = defaultPostgresPort
		}
		prompt := &survey.Input{
			Message: fmt.Sprintf("Enter DB port [%s]: ", color.New(color.Faint).Sprint(defaultPort)),
			Default: defaultPort,
		}
		survey.AskOne(prompt, &dbPort)
	}
}
