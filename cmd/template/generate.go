package template

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/briandowns/spinner"
	"github.com/golang-programming/gincli/embedded"
	"github.com/golang-programming/gincli/utils"
)

func createProjectFromTemplate(projectDir string) {
	// Initialize and start the spinner
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Suffix = " Creating project structure..."
	s.Start()
	defer s.Stop()

	// Build the root path for templates based on the template choice
	embeddedRoot := fmt.Sprintf("templates/%s", strings.ToLower(templateChoice))

	// Walk through the embedded file system
	if err := fs.WalkDir(embedded.TemplatesFS, embeddedRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Determine the relative path and target path
		relativePath, err := filepath.Rel(embeddedRoot, path)
		if err != nil {
			return err
		}
		targetPath := filepath.Join(projectDir, relativePath)

		// If it's a directory, create it
		if d.IsDir() {
			return os.MkdirAll(targetPath, os.ModePerm)
		}

		// If it's a template file, process it
		if strings.HasSuffix(d.Name(), ".tpl") {
			targetFile := strings.TrimSuffix(targetPath, ".tpl")

			// Read and process the template file
			if data, err := embedded.TemplatesFS.ReadFile(path); err == nil {
				return GenerateFileFromTemplateBytes(data, targetFile, getConfig())
			} else {
				return fmt.Errorf("failed to read embedded template file %s: %v", path, err)
			}
		}

		return nil
	}); err != nil {
		utils.LogError(fmt.Sprintf("Error while copying templates: %v", err))
	}
}

func GenerateFileFromTemplateBytes(templateData []byte, targetFile string, config map[string]string) error {
	tmpl, err := template.New("template").Parse(string(templateData))
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, config); err != nil {
		return err
	}

	// Ensure the directory exists and write the file
	if err := os.MkdirAll(filepath.Dir(targetFile), os.ModePerm); err != nil {
		return err
	}
	return os.WriteFile(targetFile, buf.Bytes(), 0644)
}

func getConfig() map[string]string {
	return map[string]string{
		"DBUsername": dbUsername,
		"DBPassword": dbPassword,
		"DBName":     dbName,
		"DBHost":     dbHost,
		"DBPort":     dbPort,
		"Module":     appName,
		"DBDriver":   strings.ToLower(dbType),
	}
}
