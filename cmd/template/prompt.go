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
		fmt.Printf("Enter your app name [%s]: ", color.New(color.Faint).Sprint(defaultAppName))
		fmt.Scanln(&appName)
		if appName == "" {
			appName = defaultAppName
		}
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
		fmt.Printf("Enter DB username [%s]: ", color.New(color.Faint).Sprint(defaultDBUsername))
		fmt.Scanln(&dbUsername)
		if dbUsername == "" {
			dbUsername = defaultDBUsername
		}
	}
	if dbPassword == "" {
		fmt.Printf("Enter DB password [%s]: ", color.New(color.Faint).Sprint(defaultDBPassword))
		fmt.Scanln(&dbPassword)
		if dbPassword == "" {
			dbPassword = defaultDBPassword
		}
	}
	if dbName == "" {
		fmt.Printf("Enter DB name [%s]: ", color.New(color.Faint).Sprint(defaultDBName))
		fmt.Scanln(&dbName)
		if dbName == "" {
			dbName = defaultDBName
		}
	}
	if dbHost == "" {
		fmt.Printf("Enter DB host [%s]: ", color.New(color.Faint).Sprint(defaultDBHost))
		fmt.Scanln(&dbHost)
		if dbHost == "" {
			dbHost = defaultDBHost
		}
	}
	if dbPort == "" {
		defaultPort := defaultMySQLPort
		if strings.ToLower(dbType) == "postgresql" {
			defaultPort = defaultPostgresPort
		}
		fmt.Printf("Enter DB port [%s]: ", color.New(color.Faint).Sprint(defaultPort))
		fmt.Scanln(&dbPort)
		if dbPort == "" {
			dbPort = defaultPort
		}
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
