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
	servicePath string
)

var ServiceCmd = &cobra.Command{
	Use:   "service <name> [path]",
	Short: "Generate a new service",
	Args:  cobra.MinimumNArgs(1), // Ensure at least one argument (service name) is provided
	Run:   createService,
}

func createService(cmd *cobra.Command, args []string) {
	serviceName = args[0]
	defaultProjectDir := filepath.Join(".", "app", "modules", serviceName, "services")

	// Check if a custom path is provided as the second argument
	if len(args) > 1 {
		servicePath = args[1]
	} else {
		servicePath = defaultProjectDir
	}

	generateServiceFile(servicePath)

	fmt.Println(color.New(color.FgGreen).Sprint("Service generated successfully at path: " + servicePath))
}
