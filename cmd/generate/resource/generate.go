package resource

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/briandowns/spinner"
	"github.com/golang-programming/gincli/embedded"
	"github.com/golang-programming/gincli/utils"
)

func generateResourceFromTemplate() {
	embeddedRoot := fmt.Sprintf("templates/others/%s/%s", strings.ToLower(transport),
		func() string {
			if createEndpoint {
				return "endpoints"
			}
			return "blank"
		}(),
	)

	s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Suffix = " Creating resource..."
	s.Start()
	defer s.Stop()

	if err := fs.WalkDir(embedded.TemplatesFS, embeddedRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(embeddedRoot, path)
		if err != nil {
			return err
		}
		targetPath := filepath.Join(resourcePath, relPath)

		if d.IsDir() {
			return os.MkdirAll(targetPath, os.ModePerm)
		}

		if strings.HasSuffix(d.Name(), ".tpl") {
			targetFile := strings.TrimSuffix(targetPath, ".tpl")
			data, err := embedded.TemplatesFS.ReadFile(path)
			if err != nil {
				return fmt.Errorf("failed to read template %s: %v", path, err)
			}

			tmpl, err := template.New("").Parse(string(data))
			if err != nil {
				return fmt.Errorf("failed to parse template %s: %v", path, err)
			}

			var buf bytes.Buffer
			if err := tmpl.Execute(&buf, getConfig()); err != nil {
				return fmt.Errorf("failed to execute template %s: %v", targetFile, err)
			}

			if err := os.WriteFile(targetFile, buf.Bytes(), 0644); err != nil {
				return fmt.Errorf("failed to write file %s: %v", targetFile, err)
			}
		}

		return nil
	}); err != nil {
		utils.LogError(fmt.Sprintf("Error while copying templates: %v", err))
	}
}

func getConfig() map[string]string {
	return map[string]string{
		"CapitalizeResourceName": utils.ToPascalCase(resourceName),
		"ResourceName":           utils.ConvertToSnakeCase(resourceName),
		"Module":                 utils.DetectModuleName(),
	}
}
