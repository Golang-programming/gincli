package template

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/golang-programming/gincli/utils"
)

func createProjectFromTemplate(templateDir, projectDir string) {
	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Suffix = " Creating project structure..."
	s.Start()
	defer s.Stop()

	err := filepath.Walk(templateDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relativePath := strings.TrimPrefix(path, templateDir)
		targetPath := filepath.Join(projectDir, relativePath)

		if info.IsDir() {
			return os.MkdirAll(targetPath, os.ModePerm)
		}

		if strings.HasSuffix(info.Name(), ".tpl") {
			targetFile := strings.TrimSuffix(targetPath, ".tpl")
			utils.GenerateFileFromTemplate(path, targetFile, nil)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error while copying templates: %v\n", err)
	}
}