// cmd/new.go
package template

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/golang-programming/gincli/utils"
	"github.com/spf13/cobra"
)

var (
	appName            string
	dbType             string
	dbHost             string
	dbName             string
	dbUsername         string
	dbPassword         string
	dbPort             string
	dbConnectionString string
	skipPrompts        bool
	templateChoice     string
)

const (
	defaultAppName        = "my-gin-app"
	defaultDBUsername     = "root"
	defaultDBType         = "MySQL"
	defaultDBPassword     = "password"
	defaultDBName         = "default"
	defaultDBHost         = "localhost"
	defaultMySQLPort      = "3306"
	defaultPostgresPort   = "5432"
	defaultTemplateChoice = "Standard"
)

var availableTemplates = []string{"standard"}

var availableDBTypes = []string{
	"MySQL",
	"PostgresQL",
}

var TemplateCmd = &cobra.Command{
	Use:   "template",
	Short: "Load application startup template",
	Run:   loadTemplate,
}

func init() {
	TemplateCmd.Flags().StringVar(&templateChoice, "template", "", fmt.Sprintf("Template: Standard (default: %s)", defaultTemplateChoice))
	TemplateCmd.Flags().StringVar(&appName, "app-name", "", fmt.Sprintf("Name of your application (default: %s)", defaultAppName))
	TemplateCmd.Flags().StringVar(&dbType, "db-type", "", "Database type: MySQL, PostgreSQL")
	TemplateCmd.Flags().StringVar(&dbConnectionString, "db-connection-string", "", "Database connection string")
	TemplateCmd.Flags().StringVar(&dbHost, "db-host", "", fmt.Sprintf("Database host (default: %s)", defaultDBHost))
	TemplateCmd.Flags().StringVar(&dbName, "db-name", "", fmt.Sprintf("Database name (default: %s)", defaultDBName))
	TemplateCmd.Flags().StringVar(&dbUsername, "db-username", "", fmt.Sprintf("Database username (default: %s)", defaultDBUsername))
	TemplateCmd.Flags().StringVar(&dbPassword, "db-password", "", fmt.Sprintf("Database password (default: %s)", defaultDBPassword))
	TemplateCmd.Flags().StringVar(&dbPort, "db-port", "", "Database port (default: 3306 for MySQL, 5432 for PostgreSQL)")
	TemplateCmd.Flags().BoolVarP(&skipPrompts, "yes", "y", false, "Skip all prompts and use default values")
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

	// Run go mod tidy with spinner
	runGoModTidy(projectDir)

	fmt.Println(color.New(color.FgGreen).Sprint("Template loaded successfully"))
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
