// ./utils/file_writer.go
package utils

import (
	"bytes"
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/AlecAivazis/survey/v2"
)

var newAppTemplates embed.FS
var otherTemplates embed.FS

func GenerateFileFromTemplate(templateName, outputPath string, data map[string]string) {
	var tmplFS embed.FS
	switch {
	case filepath.Dir(templateName) == "templates/templates":
		tmplFS = newAppTemplates
	case filepath.Dir(templateName) == "templates/new-app":
		tmplFS = newAppTemplates
	case filepath.Dir(templateName) == "templates/others":
		tmplFS = otherTemplates
	default:
		LogError(fmt.Sprintf("Unknown template directory: %s", filepath.Dir(templateName)))
	}

	tmplContent, err := tmplFS.ReadFile(templateName)
	if err != nil {
		LogError(fmt.Sprintf("Error reading embedded template: %s", err))
	}

	tmpl, err := template.New(filepath.Base(templateName)).Parse(string(tmplContent))
	if err != nil {
		LogError(fmt.Sprintf("Error parsing template: %s", err))
	}

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

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		LogError(fmt.Sprintf("Error executing template: %s", err))
	}

	err = os.WriteFile(outputPath, buf.Bytes(), 0644)
	if err != nil {
		LogError(fmt.Sprintf("Error writing file: %s", err))
	}

	LogSuccess(fmt.Sprintf("Generated file: %s", outputPath))
}
