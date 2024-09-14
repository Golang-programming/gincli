// cmd/new.go
package new

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/golang-programming/gincli/utils"
	"github.com/spf13/cobra"
)

var (
	appName            string
	dbTypeChoice       string
	dbHost             string
	dbName             string
	dbUsername         string
	dbPassword         string
	dbPort             string
	dbConnectionString string
	skipPrompts        bool
)

const (
	defaultAppName      = "my-gin-app"
	defaultDBUsername   = "root"
	defaultDBPassword   = "password"
	defaultDBName       = "default"
	defaultDBHost       = "localhost"
	defaultMySQLPort    = "3306"
	defaultPostgresPort = "5432"
)

var dbTypes = map[string]string{
	"1": "MySQL",
	"2": "PostgreSQL",
	"3": "SQLite",
	"4": "MongoDB",
}

var NewCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new Gin application with a project structure",
	Run:   createNewApp,
}

func init() {
	NewCmd.Flags().StringVar(&appName, "app-name", "", fmt.Sprintf("Name of your application (default: %s)", defaultAppName))
	NewCmd.Flags().StringVar(&dbTypeChoice, "db-type", "", "Database type: 1. MySQL, 2. PostgreSQL, 3. SQLite, 4. MongoDB (default: 1)")
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
	utils.InitializeGoModule(projectDir, appName)
	setupProjectDirectories()
	generateProjectFiles(projectDir)

	// Run go mod tidy with spinner
	runGoModTidy(projectDir)

	fmt.Println(color.New(color.FgGreen).Sprint("Application created successfully"))
	fmt.Printf("Next steps:\n")
	fmt.Printf("Go to project directory: cd %s\n", projectDir)
	fmt.Printf("Run your project: go run *.go\n")
}

func runGoModTidy(projectDir string) {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Suffix = " Running `go mod tidy`..."
	s.Start()
	defer s.Stop()

	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = projectDir
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error running `go mod tidy`: %v\n", err)
		os.Exit(1)
	}

	s.Stop()
	fmt.Println(color.New(color.FgGreen).Sprint("`go mod tidy` completed successfully."))
}

func setDefaultValues() {
	if appName == "" {
		appName = defaultAppName
	}
	if dbTypeChoice == "" {
		dbTypeChoice = "1"
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
		if dbTypeChoice == "2" {
			dbPort = defaultPostgresPort
		} else {
			dbPort = defaultMySQLPort
		}
	}
	if dbConnectionString == "" {
		dbConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)
	}
}
