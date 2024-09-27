// ./utils/detect-module.go
package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// DetectModuleName retrieves the module name from go.mod.
func DetectModuleName() string {
	goModPath, err := filepath.Abs("go.mod")
	if err != nil {
		LogError(fmt.Sprintf("Error determining the absolute path of go.mod: %s", err))
	}

	file, err := os.Open(goModPath)
	if err != nil {
		LogError(fmt.Sprintf("Error opening go.mod file: %s", err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "module") {
			moduleName := strings.TrimSpace(strings.TrimPrefix(line, "module"))
			if moduleName != "" {
				return moduleName
			}
		}
	}

	if err := scanner.Err(); err != nil {
		LogError(fmt.Sprintf("Error reading go.mod file: %s", err))
	}

	LogError("Module name not found in go.mod. Please ensure you are in the root directory of your project and that the go.mod file contains a valid module name.")
	return ""
}
