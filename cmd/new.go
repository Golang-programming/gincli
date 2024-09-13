package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang-programming/gincli/utils"
	"github.com/spf13/cobra"
)

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new Gin application with a project structure",
	Run: func(cmd *cobra.Command, args []string) {
		createNewApp()
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}

func createNewApp() {
	var appName, dbType, dbHost, dbName, dbUsername, dbPassword, dbPort string

	// Ask for app name
	fmt.Print("Enter your app name: ")
	fmt.Scanln(&appName)

	// Ask for database type
	fmt.Println("Select your database:")
	fmt.Println("1. MySQL")
	fmt.Println("2. PostgreSQL")
	fmt.Println("3. SQLite")
	fmt.Println("4. MongoDB")
	fmt.Print("Enter the number of the database type: ")
	fmt.Scanln(&dbType)

	// Collect DB Config
	fmt.Print("Enter your DB username: ")
	fmt.Scanln(&dbUsername)
	fmt.Print("Enter your DB password: ")
	fmt.Scanln(&dbPassword)
	fmt.Print("Enter your DB name: ")
	fmt.Scanln(&dbName)
	fmt.Print("Enter your DB host (e.g., localhost): ")
	fmt.Scanln(&dbHost)
	fmt.Print("Enter your DB port (e.g., 3306 for MySQL): ")
	fmt.Scanln(&dbPort)

	// Create project directory
	projectDir := filepath.Join(".", appName)
	err := os.MkdirAll(projectDir, os.ModePerm)
	if err != nil {
		fmt.Printf("Error creating project directory: %s\n", err)
		return
	}

	// Create necessary subdirectories
	os.MkdirAll(filepath.Join(projectDir, "app", "pkg", "database"), os.ModePerm)
	os.MkdirAll(filepath.Join(projectDir, "app", "controllers"), os.ModePerm)
	os.MkdirAll(filepath.Join(projectDir, "app", "services"), os.ModePerm)
	os.MkdirAll(filepath.Join(projectDir, "app", "utils"), os.ModePerm)

	// Initialize Go module
	utils.InitializeGoModule(projectDir, appName)

	// Generate all files from the structure
	generateFullStructure(appName, dbType, dbUsername, dbPassword, dbName, dbHost, dbPort, projectDir)
}

func generateFullStructure(appName, dbType, dbUsername, dbPassword, dbName, dbHost, dbPort, projectDir string) {
	// Determine DB driver
	var dbDriver string
	switch dbType {
	case "1":
		dbDriver = "mysql"
	case "2":
		dbDriver = "postgres"
	case "3":
		dbDriver = "sqlite"
	default:
		dbDriver = "mysql"
	}

	// Generate .env file
	utils.GenerateFileFromTemplate(filepath.Join("templates", "new", ".env.tpl"), filepath.Join(projectDir, ".env"), map[string]string{
		"DBUsername": dbUsername,
		"DBPassword": dbPassword,
		"DBName":     dbName,
		"DBHost":     dbHost,
		"DBPort":     dbPort,
	})

	// Generate main.go file
	utils.GenerateFileFromTemplate(filepath.Join("templates", "new", "main.go.tpl"), filepath.Join(projectDir, "main.go"), map[string]string{
		"Module": appName,
	})

	// Generate loadEnv.go file
	utils.GenerateFileFromTemplate(filepath.Join("templates", "new", "loadEnv.go.tpl"), filepath.Join(projectDir, "loadEnv.go"), map[string]string{})

	// Generate routes.go file
	utils.GenerateFileFromTemplate(filepath.Join("templates", "new", "routes.go.tpl"), filepath.Join(projectDir, "routes.go"), map[string]string{
		"Module": appName,
	})

	// Generate database.go file
	utils.GenerateFileFromTemplate(filepath.Join("templates", "new", "app", "pkg", "database", "database.go.tpl"), filepath.Join(projectDir, "app", "pkg", "database", "database.go"), map[string]string{
		"DBDriver": dbDriver,
	})

	// Generate controller.go file (Use `controller.go` instead of appending `appName`)
	utils.GenerateFileFromTemplate(filepath.Join("templates", "new", "app", "controllers", "controller.go.tpl"), filepath.Join(projectDir, "app", "controllers", "controller.go"), map[string]string{
		"Module": appName,
	})

	// Generate service.go file (Use `service.go` instead of appending `appName`)
	utils.GenerateFileFromTemplate(filepath.Join("templates", "new", "app", "services", "service.go.tpl"), filepath.Join(projectDir, "app", "services", "service.go"), map[string]string{
		"Module": appName,
	})

	// Generate sum-to-numbers.go file
	utils.GenerateFileFromTemplate(filepath.Join("templates", "new", "app", "utils", "sum-to-numbers.go.tpl"), filepath.Join(projectDir, "app", "utils", "sum_to_numbers.go"), map[string]string{})

	fmt.Println("Application created successfully with full structure!")
}
