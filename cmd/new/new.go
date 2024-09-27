// ./cmd/new/new.go
package new

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/golang-programming/gincli/utils"
	"github.com/spf13/cobra"
)

var (
	appName             string
	dbType              string
	dbHost              string
	dbName              string
	dbUsername          string
	dbPassword          string
	dbPort              string
	dbConnectionString  string
	skipPrompts         bool
	defaultAppName      = "my-gin-app"
	defaultDBType       = "MySQL"
	defaultDBUsername   = "root"
	defaultDBPassword   = "password"
	defaultDBName       = "default"
	defaultDBHost       = "localhost"
	defaultMySQLPort    = "3306"
	defaultPostgresPort = "5432"
	availableDBTypes    = []string{"MySQL", "PostgreSQL", "SQLite", "MongoDB"}
)

var NewCmd = &cobra.Command{
	Use:     "new",
	Short:   "Create a new Gin application with a project structure",
	Aliases: []string{"n", "create"},
	Run:     createNewApp,
}

func init() {
	NewCmd.Flags().StringVarP(&appName, "app-name", "a", "", fmt.Sprintf("Name of your application (default: %s)", defaultAppName))
	NewCmd.Flags().StringVarP(&dbType, "db-type", "d", "", "Database type: MySQL, PostgreSQL, SQLite, MongoDB (default: MySQL)")
	NewCmd.Flags().StringVarP(&dbConnectionString, "db-connection-string", "c", "", "Database connection string")
	NewCmd.Flags().StringVarP(&dbHost, "db-host", "H", "", fmt.Sprintf("Database host (default: %s)", defaultDBHost))
	NewCmd.Flags().StringVarP(&dbName, "db-name", "n", "", fmt.Sprintf("Database name (default: %s)", defaultDBName))
	NewCmd.Flags().StringVarP(&dbUsername, "db-username", "u", "", fmt.Sprintf("Database username (default: %s)", defaultDBUsername))
	NewCmd.Flags().StringVarP(&dbPassword, "db-password", "p", "", fmt.Sprintf("Database password (default: %s)", defaultDBPassword))
	NewCmd.Flags().StringVarP(&dbPort, "db-port", "P", "", "Database port (default: 3306 for MySQL, 5432 for PostgreSQL)")
	NewCmd.Flags().BoolVarP(&skipPrompts, "yes", "y", false, "Skip all prompts and use default values")
}

func createNewApp(cmd *cobra.Command, args []string) {
	if skipPrompts {
		setDefaultValues()
	} else {
		promptForValues()
	}

	projectDir := filepath.Join(".", appName)
	setupProjectDirectories()
	generateProjectFiles(projectDir)

	// Run go mod tidy with spinner handled inside the utility
	utils.InitializeGoModule(projectDir, appName)
	utils.RunGoModTidy(projectDir)

	utils.LogSuccess("Application created successfully")
	fmt.Println("Next steps:")
	fmt.Printf("  Go to project directory: cd %s\n", projectDir)
	fmt.Println("  Run your project: go run *.go")
}

func setDefaultValues() {
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
