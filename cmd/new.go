package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/golang-programming/gincli/utils"
	"github.com/spf13/cobra"
)

var (
	appName              string
	dbType               string
	dbHost               string
	dbName               string
	dbUsername           string
	dbPassword           string
	dbPort               string
	dbConnectionString   string
	skipPrompts          bool
	defaultAppName       = "my-gin-app"
	defaultDBType        = "1" // MySQL as default
	defaultDBHost        = "localhost"
	defaultDBName        = "testdb"
	defaultDBUsername    = "root"
	defaultDBPassword    = "password"
	defaultDBPort        = "3306"
	defaultConnectionStr = "user:password@tcp(localhost:3306)/dbname"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new Gin application with a project structure",
	Run:   createNewApp,
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Add flags for user input
	newCmd.Flags().StringVar(&appName, "app-name", "", "Name of your application")
	newCmd.Flags().StringVar(&dbType, "db-type", "", "Database type (1: MySQL, 2: PostgreSQL, 3: SQLite, 4: MongoDB)")
	newCmd.Flags().StringVar(&dbConnectionString, "db-connection-string", "", "Database connection string")
	newCmd.Flags().StringVar(&dbHost, "db-host", "", "Database host")
	newCmd.Flags().StringVar(&dbName, "db-name", "", "Database name")
	newCmd.Flags().StringVar(&dbUsername, "db-username", "", "Database username")
	newCmd.Flags().StringVar(&dbPassword, "db-password", "", "Database password")
	newCmd.Flags().StringVar(&dbPort, "db-port", "", "Database port")
	newCmd.Flags().BoolVarP(&skipPrompts, "yes", "y", false, "Skip all prompts and use default values")
}

func createNewApp(cmd *cobra.Command, args []string) {
	if skipPrompts {
		applyDefaultValues()
	} else {
		askForMissingValues()
	}

	projectDir := filepath.Join(".", appName)
	createDirectoriesFromTemplate("templates/new/app", projectDir)

	utils.InitializeGoModule(projectDir, appName)
	generateProjectFiles(appName, dbType, getDBConfig(), projectDir)

	fmt.Println("Application created successfully with full structure and Docker files!")
}

// Use default values if --yes or -y flag is passed
func applyDefaultValues() {
	appName = defaultAppName
	dbType = defaultDBType
	dbHost = defaultDBHost
	dbName = defaultDBName
	dbUsername = defaultDBUsername
	dbPassword = defaultDBPassword
	dbPort = defaultDBPort
	dbConnectionString = defaultConnectionStr
}

// Ask for missing values if no --yes flag is passed
func askForMissingValues() {
	if appName == "" {
		fmt.Print("Enter your app name: ")
		fmt.Scanln(&appName)
	}
	if dbType == "" {
		fmt.Println("Select your database: 1. MySQL 2. PostgreSQL 3. SQLite 4. MongoDB")
		fmt.Scanln(&dbType)
	}
	if dbConnectionString == "" {
		if dbUsername == "" {
			fmt.Print("Enter DB username: ")
			fmt.Scanln(&dbUsername)
		}
		if dbPassword == "" {
			fmt.Print("Enter DB password: ")
			fmt.Scanln(&dbPassword)
		}
		if dbName == "" {
			fmt.Print("Enter DB name: ")
			fmt.Scanln(&dbName)
		}
		if dbHost == "" {
			fmt.Print("Enter DB host (e.g., localhost): ")
			fmt.Scanln(&dbHost)
		}
		if dbPort == "" {
			fmt.Print("Enter DB port (e.g., 3306 for MySQL): ")
			fmt.Scanln(&dbPort)
		}
	} else {
		// If the user passes a connection string, parse the individual components
		parseConnectionString(dbConnectionString)
	}
}

// Function to parse the DB connection string and set the variables
func parseConnectionString(connectionString string) {
	// You can add logic to parse the connection string and split it into dbHost, dbUsername, dbPassword, etc.
	parts := strings.Split(connectionString, "@")
	if len(parts) == 2 {
		dbCredentials := strings.Split(parts[0], ":")
		if len(dbCredentials) == 2 {
			dbUsername = dbCredentials[0]
			dbPassword = dbCredentials[1]
		}
		dbHostAndPort := strings.Split(parts[1], "/")
		hostPort := strings.Split(dbHostAndPort[0], ":")
		if len(hostPort) == 2 {
			dbHost = hostPort[0]
			dbPort = hostPort[1]
		}
		dbName = dbHostAndPort[1]
	}
}

// Return the DB configuration
func getDBConfig() map[string]string {
	return map[string]string{
		"DBUsername": dbUsername,
		"DBPassword": dbPassword,
		"DBName":     dbName,
		"DBHost":     dbHost,
		"DBPort":     dbPort,
	}
}

func createDirectoriesFromTemplate(templateDir, projectDir string) {
	err := filepath.Walk(templateDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip root template directory
		if templateDir == path {
			return nil
		}

		// Replace "templates/new/app" with the actual project directory structure
		relativePath := strings.TrimPrefix(path, templateDir)
		targetPath := filepath.Join(projectDir, relativePath)

		// If it's a directory, create it
		if info.IsDir() {
			return os.MkdirAll(targetPath, os.ModePerm)
		}

		// If it's a file, prepare it for template processing
		if strings.HasSuffix(info.Name(), ".tpl") {
			// Copy the template file to the corresponding path without the ".tpl" suffix
			targetFile := strings.TrimSuffix(targetPath, ".tpl")
			utils.GenerateFileFromTemplate(path, targetFile, nil)
			return nil
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error while copying templates: %v\n", err)
	}
}

func generateProjectFiles(appName, dbType string, dbConfig map[string]string, projectDir string) {
	dbDriver := map[string]string{"1": "mysql", "2": "postgres", "3": "sqlite"}[dbType]
	if dbDriver == "" {
		dbDriver = "mysql" // Default to MySQL
	}

	// Files that require template variables
	templates := map[string]string{
		"templates/new/.env.tpl":                      ".env",
		"templates/new/main.go.tpl":                   "main.go",
		"templates/new/loadEnv.go.tpl":                "loadEnv.go",
		"templates/new/routes.go.tpl":                 "routes.go",
		"templates/new/pkg/database/database.go.tpl":  "pkg/database/database.go",
		"templates/new/controllers/controller.go.tpl": "controllers/" + appName + "_controller.go",
		"templates/new/services/service.go.tpl":       "services/" + appName + "_service.go",
		"templates/new/utils/sum-to-numbers.go.tpl":   "utils/sum_to_numbers.go",
		"templates/new/Dockerfile.tpl":                "Dockerfile",
	}

	for tpl, output := range templates {
		config := dbConfig
		if tpl == "templates/new/pkg/database/database.go.tpl" {
			config = map[string]string{"DBDriver": dbDriver}
		} else if tpl == "templates/new/main.go.tpl" || tpl == "templates/new/routes.go.tpl" {
			config = map[string]string{"Module": appName}
		}
		utils.GenerateFileFromTemplate(filepath.Join(tpl), filepath.Join(projectDir, output), config)
	}

	// Generate docker-compose.yml if MySQL or PostgreSQL is selected
	if dbType == "1" || dbType == "2" {
		utils.GenerateFileFromTemplate(filepath.Join("templates/new/docker-compose.yml.tpl"), filepath.Join(projectDir, "docker-compose.yml"), map[string]string{"AppName": appName})
	}
}
