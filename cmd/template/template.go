// ./cmd/template/template.go
package template

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang-programming/gincli/utils"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

var (
	appName               string
	dbType                string
	dbHost                string
	dbName                string
	dbUsername            string
	dbPassword            string
	dbPort                string
	dbConnectionString    string
	skipPrompts           bool
	templateChoice        string
	defaultAppName        = "my-gin-app"
	defaultDBUsername     = "root"
	defaultDBType         = "MySQL"
	defaultDBPassword     = "password"
	defaultDBName         = "default"
	defaultDBHost         = "localhost"
	defaultMySQLPort      = "3306"
	defaultPostgresPort   = "5432"
	defaultTemplateChoice = "Standard"
	availableTemplates    = []string{"standard"}
	availableDBTypes      = []string{"MySQL", "PostgreSQL"}
)

var TemplateCmd = &cobra.Command{
	Use:   "template",
	Short: "Load application startup template",
	Run:   loadTemplate,
}

func init() {
	TemplateCmd.Flags().StringVarP(&templateChoice, "template", "t", "", fmt.Sprintf("Template: Standard (default: %s)", defaultTemplateChoice))
	TemplateCmd.Flags().StringVarP(&appName, "app-name", "a", "", fmt.Sprintf("Name of your application (default: %s)", defaultAppName))
	TemplateCmd.Flags().StringVarP(&dbType, "db-type", "d", "", "Database type: MySQL, PostgreSQL")
	TemplateCmd.Flags().StringVarP(&dbConnectionString, "db-connection-string", "c", "", "Database connection string")
	TemplateCmd.Flags().StringVarP(&dbHost, "db-host", "H", "", fmt.Sprintf("Database host (default: %s)", defaultDBHost))
	TemplateCmd.Flags().StringVarP(&dbName, "db-name", "n", "", fmt.Sprintf("Database name (default: %s)", defaultDBName))
	TemplateCmd.Flags().StringVarP(&dbUsername, "db-username", "u", "", fmt.Sprintf("Database username (default: %s)", defaultDBUsername))
	TemplateCmd.Flags().StringVarP(&dbPassword, "db-password", "p", "", fmt.Sprintf("Database password (default: %s)", defaultDBPassword))
	TemplateCmd.Flags().StringVarP(&dbPort, "db-port", "P", "", "Database port (default: 3306 for MySQL, 5432 for PostgreSQL)")
	TemplateCmd.Flags().BoolVarP(&skipPrompts, "yes", "y", false, "Skip all prompts and use default values")
}

// CustomTemplateHelpFunc defines a custom help function for the template command.
func CustomTemplateHelpFunc(cmd *cobra.Command, args []string) {
	utils.LogInfo("Template Command - Load application startup template.\n")
	utils.LogInfo("Usage:")
	utils.LogInfo("  gincli template [flags]\n")
	utils.LogInfo("Available Flags:")

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Flag", "Shorthand", "Description", "Default"})

	flags := []struct {
		Name        string
		Shorthand   string
		Description string
		Default     string
	}{
		{"template", "t", "Template to use (Standard)", defaultTemplateChoice},
		{"app-name", "a", "Name of your application", defaultAppName},
		{"db-type", "d", "Database type (MySQL, PostgreSQL)", defaultDBType},
		{"db-connection-string", "c", "Database connection string", ""},
		{"db-host", "H", "Database host", defaultDBHost},
		{"db-name", "n", "Database name", defaultDBName},
		{"db-username", "u", "Database username", defaultDBUsername},
		{"db-password", "p", "Database password", defaultDBPassword},
		{"db-port", "P", "Database port", "3306 or 5432"},
		{"yes", "y", "Skip prompts and use default values", "false"},
	}

	for _, flag := range flags {
		table.Append([]string{fmt.Sprintf("--%s", flag.Name), fmt.Sprintf("-%s", flag.Shorthand), flag.Description, flag.Default})
	}

	table.Render()

	utils.LogInfo("\nUse \"gincli template [flags]\" to customize the template loading.")
}

func loadTemplate(cmd *cobra.Command, args []string) {
	if skipPrompts {
		setDefaultValues()
	} else {
		promptForValues()
	}

	projectDir := filepath.Join(".", appName)
	createProjectFromTemplate(fmt.Sprintf("templates/templates/%s", strings.ToLower(templateChoice)), projectDir)

	utils.InitializeGoModule(projectDir, appName)

	utils.RunGoModTidy(projectDir)

	utils.LogSuccess("Template loaded successfully")
	fmt.Println("Next steps:")
	fmt.Printf("  Go to project directory: cd %s\n", projectDir)
	fmt.Println("  Run your project: go run *.go")
}

func setDefaultValues() {
	if templateChoice == "" {
		templateChoice = defaultTemplateChoice
	}
	if appName == "" {
		appName = defaultAppName
	}
	if dbType == "" {
		dbType = defaultDBType
	}
	if dbHost == "" {
		dbHost = defaultDBHost
	}
	if dbName == "" {
		dbName = defaultDBName
	}
	if dbUsername == "" {
		dbUsername = defaultDBUsername
	}
	if dbPassword == "" {
		dbPassword = defaultDBPassword
	}
	if dbPort == "" {
		if strings.ToLower(dbType) == "postgresql" {
			dbPort = defaultPostgresPort
		} else {
			dbPort = defaultMySQLPort
		}
	}
	if dbConnectionString == "" {
		dbConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)
	}
}
