// ./utils/file_writer.go
package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/AlecAivazis/survey/v2"
	"github.com/golang-programming/gincli/embedded" // Adjust the import path based on your module
)

// GenerateFileFromTemplate generates a file from a template using embedded templates.
func GenerateFileFromTemplate(templatePath, outputPath string, data map[string]string) {
	// Check if the file already exists
	if _, err := os.Stat(outputPath); err == nil {
		// File exists, prompt for overwrite
		overwrite := false
		prompt := &survey.Confirm{
			Message: fmt.Sprintf("File %s already exists. Do you want to overwrite it?", outputPath),
			Default: false,
		}
		survey.AskOne(prompt, &overwrite)
		if !overwrite {
			LogWarning(fmt.Sprintf("Skipped generating file: %s", outputPath))
			return
		}
	}

	// Create the directories if they do not exist
	if err := CreateDirectories(filepath.Dir(outputPath)); err != nil {
		LogError(fmt.Sprintf("Error creating directories: %s", err))
	}

	// Read the template from the embedded filesystem
	tmplBytes, err := embedded.TemplatesFS.ReadFile(templatePath)
	if err != nil {
		LogError(fmt.Sprintf("Error reading embedded template (%s): %s", templatePath, err))
	}

	// Parse the template
	tmpl, err := template.New(filepath.Base(templatePath)).Parse(string(tmplBytes))
	if err != nil {
		LogError(fmt.Sprintf("Error parsing template (%s): %s", templatePath, err))
	}

	// Create or truncate the output file
	outputFile, err := os.Create(outputPath)
	if err != nil {
		LogError(fmt.Sprintf("Error creating file (%s): %s", outputPath, err))
	}
	defer outputFile.Close()

	// Execute the template with provided data
	err = tmpl.Execute(outputFile, data)
	if err != nil {
		LogError(fmt.Sprintf("Error executing template (%s): %s", templatePath, err))
	}

	LogSuccess(fmt.Sprintf("Generated file: %s", outputPath))
}
