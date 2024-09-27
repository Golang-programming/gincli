package utils

import (
	"fmt"
	"os/exec"
)

func InitializeGoModule(projectDir, appName string) {
	cmd := exec.Command("go", "mod", "init", appName)
	cmd.Dir = projectDir // Set the directory where the command should run
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error initializing Go module: %s\n", string(output))
		return
	}

	fmt.Println("Go module initialized successfully!")
}
