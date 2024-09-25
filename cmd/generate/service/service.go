// cmd/generate/service/service.go
package service

import (
	"fmt"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/golang-programming/gincli/utils"
	"github.com/spf13/cobra"
)

var (
	serviceName string
	skipPrompts bool
)

var ServiceCmd = &cobra.Command{
	Use:   "service <name>",
	Short: "Generate a new service",
	Run:   createService,
}

func init() {
	ServiceCmd.Flags().BoolVarP(&skipPrompts, "yes", "y", false, "Skip prompts and use default values")
}

func createService(cmd *cobra.Command, args []string) {
	serviceName = args[0]

	promptForValues()

	projectDir := filepath.Join(".", "app", "services")
	utils.CreateDirectories([]string{projectDir})

	generateServiceFile(projectDir)

	fmt.Println(color.New(color.FgGreen).Sprint("Service generated successfully"))
}
