package new

import (
	"path/filepath"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/golang-programming/gincli/utils"
)

func generateProjectFiles(projectDir string) {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Suffix = " Generating project files..."
	s.Start()
	defer s.Stop()

	templates := map[string]string{
		"templates/new-app/.env.tpl":                         ".env",
		"templates/new-app/loadEnv.go.tpl":                   "loadEnv.go",
		"templates/new-app/app/service/service.go.tpl":       "app/service/service.go",
		"templates/new-app/app/utils/sum-to-numbers.go.tpl":  "app/utils/sum-to-numbers.go",
		"templates/new-app/Dockerfile.tpl":                   "Dockerfile",
		"templates/new-app/main.go.tpl":                      "main.go",
		"templates/new-app/routes.go.tpl":                    "routes.go",
		"templates/new-app/app/controller/controller.go.tpl": "app/controller/controller.go",
	}

	if strings.ToLower(dbType) == "mongodb" {
		templates["templates/new-app/app/pkg/database/mongodb-database.go.tpl"] = "app/pkg/database/database.go"
	} else if strings.ToLower(dbType) != "sqlite" {
		templates["templates/new-app/app/pkg/database/database.go.tpl"] = "app/pkg/database/database.go"
		templates["templates/new-app/docker-compose.yml.tpl"] = "docker-compose.yml"
	} else {
		templates["templates/new-app/app/pkg/database/sqlite-database.go.tpl"] = "app/pkg/database/database.go"
	}

	for tpl, output := range templates {
		utils.GenerateFileFromTemplate(tpl, filepath.Join(projectDir, output), getConfig())
	}

}

func getConfig() map[string]string {
	dbDriver := strings.ToLower(dbType)

	return map[string]string{
		"DBUsername": dbUsername,
		"DBPassword": dbPassword,
		"DBName":     dbName,
		"DBHost":     dbHost,
		"DBPort":     dbPort,
		"AppName":    appName,
		"Module":     appName,
		"DBDriver":   dbDriver,
		"DBUri":      mongodbUri,
	}
}
