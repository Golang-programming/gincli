// ./utils/file_writer.go
package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/AlecAivazis/survey/v2"
)

// GenerateFileFromTemplate generates a file from a template using explicit paths.
// It prompts the user before overwriting existing files.
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

	// Parse the template
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		LogError(fmt.Sprintf("Error parsing template: %s", err))
	}

	// Create or truncate the output file
	outputFile, err := os.Create(outputPath)
	if err != nil {
		LogError(fmt.Sprintf("Error creating file: %s", err))
	}
	defer outputFile.Close()

	// Execute the template with provided data
	err = tmpl.Execute(outputFile, data)
	if err != nil {
		LogError(fmt.Sprintf("Error executing template: %s", err))
	}

	LogSuccess(fmt.Sprintf("Generated file: %s", outputPath))
}
