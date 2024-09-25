package template

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/fatih/color"
)

func promptForValues() {
	promptForAppName()
	promptForTemplate()
	promptForDBType()
	if dbConnectionString == "" {
		promptForDBConfig()
		dbConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)
	} else {
		parseConnectionString(dbConnectionString)
	}
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

func promptForAppName() {
	if appName == "" {
		prompt := &survey.Input{
			Message: fmt.Sprintf("Enter your app name [%s]: ", color.New(color.Faint).Sprint(defaultAppName)),
			Default: defaultAppName,
		}
		survey.AskOne(prompt, &appName)
	}
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
		prompt := &survey.Input{
			Message: fmt.Sprintf("Enter DB password [%s]: ", color.New(color.Faint).Sprint(defaultDBPassword)),
			Default: defaultDBPassword,
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
		if strings.ToLower(dbType) == "postgresql" {
			defaultPort = defaultPostgresPort
		}
		prompt := &survey.Input{
			Message: fmt.Sprintf("Enter DB port [%s]: ", color.New(color.Faint).Sprint(defaultPort)),
			Default: defaultPort,
		}
		survey.AskOne(prompt, &dbPort)
	}
}

func parseConnectionString(connectionString string) {
	parts := strings.Split(connectionString, "@")
	if len(parts) == 2 {
		dbCredentials := strings.Split(parts[0], ":")
		if len(dbCredentials) == 2 {
			dbUsername = dbCredentials[0]
			dbPassword = dbCredentials[1]
		}
		dbHostAndPort := strings.Split(parts[1], "/")
		hostPort := strings.TrimPrefix(dbHostAndPort[0], "tcp(")
		hostPort = strings.TrimSuffix(hostPort, ")")
		hostPortParts := strings.Split(hostPort, ":")
		if len(hostPortParts) == 2 {
			dbHost = hostPortParts[0]
			dbPort = hostPortParts[1]
		}
		if len(dbHostAndPort) > 1 {
			dbName = dbHostAndPort[1]
		}
	}
}
