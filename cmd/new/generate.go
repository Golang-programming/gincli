package new

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/golang-programming/gincli/utils"
)

func createProjectFromTemplate(templateDir, projectDir string) {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Suffix = " Creating project structure..."
	s.Start()
	defer s.Stop()

	err := filepath.Walk(templateDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relativePath := strings.TrimPrefix(path, templateDir)
		targetPath := filepath.Join(projectDir, relativePath)

		if info.IsDir() {
			return os.MkdirAll(targetPath, os.ModePerm)
		}

		if strings.HasSuffix(info.Name(), ".tpl") {
			targetFile := strings.TrimSuffix(targetPath, ".tpl")
			utils.GenerateFileFromTemplate(path, targetFile, nil)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error while copying templates: %v\n", err)
	}
}

func generateProjectFiles(appName string, dbTypeChoice string, dbConfig map[string]string, projectDir string) {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Suffix = " Generating project files..."
	s.Start()
	defer s.Stop()

	dbDriver := strings.ToLower(dbTypes[dbTypeChoice])
	if dbDriver == "" {
		dbDriver = "mysql"
	}

	templates := map[string]string{
		"templates/new/.env.tpl":                         ".env",
		"templates/new/main.go.tpl":                      "main.go",
		"templates/new/loadEnv.go.tpl":                   "loadEnv.go",
		"templates/new/routes.go.tpl":                    "routes.go",
		"templates/new/app/pkg/database/database.go.tpl": "app/pkg/database/database.go",
		"templates/new/app/controller/controller.go.tpl": "app/controller/controller.go",
		"templates/new/app/service/service.go.tpl":       "app/service/service.go",
		"templates/new/app/utils/sum-to-numbers.go.tpl":  "app/utils/sum-to-numbers.go",
		"templates/new/Dockerfile.tpl":                   "Dockerfile",
	}

	for tpl, output := range templates {
		config := dbConfig

		config["Module"] = appName
		config["DBDriver"] = dbDriver

		utils.GenerateFileFromTemplate(tpl, filepath.Join(projectDir, output), config)
	}

	if dbTypeChoice == "1" || dbTypeChoice == "2" {
		utils.GenerateFileFromTemplate("templates/new/docker-compose.yml.tpl", filepath.Join(projectDir, "docker-compose.yml"), map[string]string{"AppName": appName})
	}
}

func getDBConfig() map[string]string {
	return map[string]string{
		"DBUsername": dbUsername,
		"DBPassword": dbPassword,
		"DBName":     dbName,
		"DBHost":     dbHost,
		"DBPort":     dbPort,
	}
}
