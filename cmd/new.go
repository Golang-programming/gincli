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
	// Define flags for shorthand properties
	dbTypeFlag     string
	dbUsernameFlag string
	dbPasswordFlag string
	dbNameFlag     string
	dbHostFlag     string
	dbPortFlag     string
	connStringFlag string
	yesFlag        bool
)

var newCmd = &cobra.Command{
	Use:   "new [app_name]",
	Short: "Create a new Gin application with a project structure",
	Args:  cobra.ExactArgs(1), // App name is mandatory
	Run:   createNewApp,
}

func init() {
	// Define flags for the command
	newCmd.Flags().StringVarP(&dbTypeFlag, "db-type", "d", "", "Specify the database type (mysql, postgres, sqlite, mongodb)")
	newCmd.Flags().StringVarP(&dbUsernameFlag, "db-username", "u", "", "Specify the database username")
	newCmd.Flags().StringVarP(&dbPasswordFlag, "db-password", "p", "", "Specify the database password")
	newCmd.Flags().StringVarP(&dbNameFlag, "db-name", "n", "", "Specify the database name")
	newCmd.Flags().StringVarP(&dbHostFlag, "db-host", "h", "localhost", "Specify the database host (default: localhost)")
	newCmd.Flags().StringVarP(&dbPortFlag, "db-port", "P", "", "Specify the database port (e.g., 3306 for MySQL)")
	newCmd.Flags().StringVar(&connStringFlag, "conn-string", "", "Pass a database connection string instead of individual parameters")
	newCmd.Flags().BoolVarP(&yesFlag, "yes", "y", false, "Use default values for all options")
	rootCmd.AddCommand(newCmd)
}

func createNewApp(cmd *cobra.Command, args []string) {
	appName := args[0] // Get the app name from the command arguments

	// Gather inputs
	dbType, dbConfig := gatherInputs()

	// Create project directory and set up structure
	projectDir := filepath.Join(".", appName)
	createDirectoriesFromTemplate("templates/new/app", projectDir)

	// Initialize Go module
	utils.InitializeGoModule(projectDir, appName)

	// Generate all project files from templates
	generateProjectFiles(appName, dbType, dbConfig, projectDir)

	fmt.Println("Application created successfully with full structure and Docker files!")
}

func gatherInputs() (string, map[string]string) {
	if yesFlag {
		// Use default values if the -y flag is provided
		return "mysql", map[string]string{
			"DBUsername": "root",
			"DBPassword": "password",
			"DBName":     "mydb",
			"DBHost":     "localhost",
			"DBPort":     "3306",
		}
	}

	// Check for connection string
	if connStringFlag != "" {
		dbType, dbConfig := parseConnectionString(connStringFlag)
		return dbType, dbConfig
	}

	// Prompt for missing variables
	if dbTypeFlag == "" {
		fmt.Println("Select your database: 1. MySQL, 2. PostgreSQL, 3. SQLite, 4. MongoDB")
		fmt.Scanln(&dbTypeFlag)
	}

	dbConfig := map[string]string{
		"DBUsername": getOrPrompt(dbUsernameFlag, "Enter DB username"),
		"DBPassword": getOrPrompt(dbPasswordFlag, "Enter DB password"),
		"DBName":     getOrPrompt(dbNameFlag, "Enter DB name"),
		"DBHost":     getOrPrompt(dbHostFlag, "Enter DB host (e.g., localhost)"),
		"DBPort":     getOrPrompt(dbPortFlag, "Enter DB port (e.g., 3306 for MySQL)"),
	}

	return dbTypeFlag, dbConfig
}

// If the value is missing, ask the user for input
func getOrPrompt(value, prompt string) string {
	if value == "" {
		fmt.Print(prompt + ": ")
		fmt.Scanln(&value)
	}
	return value
}

// Function to parse a database connection string
func parseConnectionString(connStr string) (string, map[string]string) {
	// Example: "user:password@tcp(localhost:3306)/dbname"
	config := map[string]string{}
	parts := strings.Split(connStr, "@")
	credentials := strings.Split(parts[0], ":")
	config["DBUsername"] = credentials[0]
	config["DBPassword"] = credentials[1]

	// Parse host and database
	hostAndDB := strings.Split(parts[1], "/")
	config["DBHost"] = strings.Split(hostAndDB[0], ":")[0]
	config["DBPort"] = strings.Split(hostAndDB[0], ":")[1]
	config["DBName"] = hostAndDB[1]

	return "mysql", config // Assuming MySQL, modify to infer DB type if needed
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

		// If it's a file, process it as a template without expecting a return value
		if strings.HasSuffix(info.Name(), ".tpl") {
			// Copy the template file to the corresponding path without the ".tpl" suffix
			targetFile := strings.TrimSuffix(targetPath, ".tpl")
			utils.GenerateFileFromTemplate(path, targetFile, nil)
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
