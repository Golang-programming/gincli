package utils

import (
	"fmt"
	"os"
	"os/exec"
	"text/template"
)

// InitializeGoModule initializes the Go module without changing the working directory
func InitializeGoModule(projectDir, appName string) {
	// Instead of changing the working directory, use the full path for go mod init
	cmd := exec.Command("go", "mod", "init", appName)
	cmd.Dir = projectDir // Set the directory where the command should run
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error initializing Go module: %s\n", string(output))
		return
	}

	fmt.Println("Go module initialized successfully!")
}

// GenerateFileFromTemplate generates a file from a template using explicit paths
func GenerateFileFromTemplate(templatePath, outputPath string, data map[string]string) {
	// Print the paths for debugging
	// cwd, _ := os.Getwd()
	// fmt.Println("Current working directory:", cwd)
	// fmt.Println("Looking for template at:", templatePath)
	// fmt.Println("Generating file at:", outputPath)

	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		fmt.Printf("Error parsing template: %s\n", err)
		return
	}

	outputFile, err := os.Create(outputPath)
	if err != nil {
		fmt.Printf("Error creating file: %s\n", err)
		return
	}
	defer outputFile.Close()

	err = tmpl.Execute(outputFile, data)
	if err != nil {
		fmt.Printf("Error executing template: %s\n", err)
		return
	}

	fmt.Printf("Generated file: %s\n", outputPath)
}
