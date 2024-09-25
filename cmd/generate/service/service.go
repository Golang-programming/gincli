// cmd/generate/service/service.go
package service

import (
	"fmt"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	serviceName string
)

var ServiceCmd = &cobra.Command{
	Use:   "service <name>",
	Short: "Generate a new service",
	Run:   createService,
}

func createService(cmd *cobra.Command, args []string) {
	serviceName = args[0]

	promptForValues()

	projectDir := filepath.Join(".", "app", "modules", serviceName, "services")

	generateServiceFile(projectDir)

	fmt.Println(color.New(color.FgGreen).Sprint("Service generated successfully"))
}
