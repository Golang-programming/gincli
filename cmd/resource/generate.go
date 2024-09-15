package resource

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/golang-programming/gincli/utils"
)

func createResourceFromTemplate() {
	templatePath := fmt.Sprintf("templates/others/resource/%s", transport)
	resourcePath := fmt.Sprintf("app/modules/%s", resourceName)

	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Suffix = " Creating resource..."
	s.Start()
	defer s.Stop()

	err := filepath.Walk(templatePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relativePath := strings.TrimPrefix(path, templatePath)
		targetPath := filepath.Join(resourcePath, relativePath)

		if info.IsDir() {
			return os.MkdirAll(targetPath, os.ModePerm)
		}

		if strings.HasSuffix(info.Name(), ".tpl") {
			targetFile := strings.TrimSuffix(targetPath, ".tpl")
			utils.GenerateFileFromTemplate(path, targetFile, getConfig())
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error while copying templates: %v\n", err)
	}
}

func getConfig() map[string]string {
	return map[string]string{
		"CapitalizeResourceName": utils.Capitalize(resourceName),
		"ResourceName":           strings.ToLower(resourceName),
		"Module":                 utils.DetectModuleName(),
	}
}
