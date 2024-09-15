package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func DetectModuleName() string {
	goModPath, err := filepath.Abs("go.mod")
	fmt.Println("goModPath", goModPath)
	if err != nil {
		log.Fatal("Error determining the absolute path of go.mod: ", err)
	}

	file, err := os.Open(goModPath)
	if err != nil {
		log.Fatalf("Error opening go.mod file. Please ensure you are in the root directory of your project. Error: %v", err)
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
		log.Fatalf("Error reading go.mod file: %v", err)
	}

	log.Fatal("Module name not found in go.mod. Please ensure you are in the root directory of your project and that the go.mod file contains a valid module name.")
	return ""
}
