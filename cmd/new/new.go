// cmd/new.go
package new

import (
	"fmt"
	"path/filepath"

	"github.com/fatih/color"
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
	availableDBTypes    = []string{"MySQL", "PostgresQL", "SQlite"}
)

var NewCmd = &cobra.Command{
	Use:     "new",
	Short:   "Create a new Gin application with a project structure",
	Aliases: []string{"n", "create"},
	Run:     createNewApp,
}

func init() {
	NewCmd.Flags().StringVar(&appName, "app-name", "", fmt.Sprintf("Name of your application (default: %s)", defaultAppName))
	NewCmd.Flags().StringVar(&dbType, "db-type", "", "Database type: MySQL, PostgreSQL, SQlite (default: MySQL)")
	NewCmd.Flags().StringVar(&dbConnectionString, "db-connection-string", "", "Database connection string")
	NewCmd.Flags().StringVar(&dbHost, "db-host", "", fmt.Sprintf("Database host (default: %s)", defaultDBHost))
	NewCmd.Flags().StringVar(&dbName, "db-name", "", fmt.Sprintf("Database name (default: %s)", defaultDBName))
	NewCmd.Flags().StringVar(&dbUsername, "db-username", "", fmt.Sprintf("Database username (default: %s)", defaultDBUsername))
	NewCmd.Flags().StringVar(&dbPassword, "db-password", "", fmt.Sprintf("Database password (default: %s)", defaultDBPassword))
	NewCmd.Flags().StringVar(&dbPort, "db-port", "", "Database port (default: 3306 for MySQL, 5432 for PostgreSQL)")
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

	// Run go mod tidy with spinner
	utils.InitializeGoModule(projectDir, appName)
	utils.RunGoModTidy(projectDir)

	fmt.Println(color.New(color.FgGreen).Sprint("Application created successfully"))
	fmt.Printf("Next steps:\n")
	fmt.Printf("Go to project directory: cd %s\n", projectDir)
	fmt.Printf("Run your project: go run *.go\n")
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
		if dbType == defaultDBType {
			dbPort = defaultPostgresPort
		} else {
			dbPort = defaultMySQLPort
		}
	}
	if dbConnectionString == "" {
		dbConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)
	}
}
