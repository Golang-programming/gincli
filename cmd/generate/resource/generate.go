// ./cmd/generate/resource/generate.go
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

func generateResourceFromTemplate() {
	templatePath := fmt.Sprintf("templates/others/%s/endpoints", strings.ToLower(transport))
	resource := resourcePath + "/" + resourceName
	if !createEndpoint {
		templatePath = fmt.Sprintf("templates/others/%s/blank", strings.ToLower(transport))
	}

	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Suffix = " Creating resource..."
	s.Start()
	defer s.Stop()

	err := filepath.Walk(templatePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relativePath := strings.TrimPrefix(path, templatePath)
		targetPath := filepath.Join(resource, relativePath)

		if strings.HasSuffix(info.Name(), ".tpl") {
			targetFile := strings.TrimSuffix(targetPath, ".tpl")
			fmt.Print(path, targetFile)
			utils.GenerateFileFromTemplate(path, targetFile, getConfig())
		}

		return nil
	})

	if err != nil {
		utils.LogError(fmt.Sprintf("Error while copying templates: %v", err))
	}
}

func getConfig() map[string]string {
	return map[string]string{
		"CapitalizeResourceName": utils.Capitalize(resourceName),
		"ResourceName":           utils.ConvertToSnakeCase(resourceName),
		"Module":                 utils.DetectModuleName(),
	}
}
