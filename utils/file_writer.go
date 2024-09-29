package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/AlecAivazis/survey/v2"
	"github.com/golang-programming/gincli/embedded"
)

func GenerateFileFromTemplate(templatePath, outputPath string, data map[string]string) {
	if _, err := os.Stat(outputPath); err == nil {
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

	if err := CreateDirectories(filepath.Dir(outputPath)); err != nil {
		LogError(fmt.Sprintf("Error creating directories: %s", err))
	}

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
