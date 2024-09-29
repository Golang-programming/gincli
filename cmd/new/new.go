// ./cmd/new/new.go
package new

import (
	"fmt"
	"path/filepath"

	"github.com/golang-programming/gincli/utils"
	"github.com/spf13/cobra"
)

var (
	appName             string
	mongodbUri          string
	dbType              string
	dbHost              string
	dbName              string
	dbUsername          string
	dbPassword          string
	dbPort              string
	skipPrompts         bool
	defaultAppName      = "my-gin-app"
	defaultDBType       = "MySQL"
	defaultDBUsername   = "root"
	defaultDBPassword   = "password"
	defaultDBName       = "default"
	defaultDBHost       = "localhost"
	defaultMySQLPort    = "3306"
	defaultPostgresPort = "5432"
	availableDBTypes    = []string{"MySQL", "Postgres", "SQLite", "MongoDB"}
)

var NewCmd = &cobra.Command{
	Use:     "new <name>",
	Short:   "Create a new Gin application with a project structure",
	Aliases: []string{"n", "create"},
	Args:    cobra.MinimumNArgs(1),
	Run:     createNewApp,
}

func init() {
	NewCmd.Flags().StringVarP(&dbType, "db-type", "d", "", "Database type: MySQL, Postgres, SQLite, MongoDB (default: MySQL)")
	NewCmd.Flags().StringVarP(&dbHost, "db-host", "H", "", fmt.Sprintf("Database host (default: %s)", defaultDBHost))
	NewCmd.Flags().StringVarP(&dbName, "db-name", "n", "", fmt.Sprintf("Database name (default: %s)", defaultDBName))
	NewCmd.Flags().StringVarP(&dbUsername, "db-username", "u", "", fmt.Sprintf("Database username (default: %s)", defaultDBUsername))
	NewCmd.Flags().StringVarP(&dbPassword, "db-password", "p", "", fmt.Sprintf("Database password (default: %s)", defaultDBPassword))
	NewCmd.Flags().StringVarP(&dbPort, "db-port", "P", "", "Database port (default: 3306 for MySQL, 5432 for Postgres)")
	NewCmd.Flags().BoolVarP(&skipPrompts, "yes", "y", false, "Skip all prompts and use default values")
}

func createNewApp(cmd *cobra.Command, args []string) {
	appName = args[0]

	if skipPrompts {
		setDefaultValues()
	} else {
		promptForValues()
	}

	projectDir := filepath.Join(".", appName)
	generateProjectFiles(projectDir)

	utils.InitializeGoModule(projectDir, appName)
	utils.RunGoModTidy(projectDir)

	utils.LogSuccess("Application created successfully")
	fmt.Println("Next steps:")
	fmt.Printf("  Go to project directory: cd %s\n", projectDir)
	fmt.Println("  Run your project: go run *.go")
}

func setDefaultValues() {
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
		dbPort = defaultMySQLPort
	}
}
