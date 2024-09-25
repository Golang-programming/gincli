// ./cmd/generate/resource/resource.go
package resource

import (
	"path/filepath"

	"github.com/golang-programming/gincli/utils"
	"github.com/spf13/cobra"
)

var (
	resourceName   string
	resourcePath   string
	transport      string
	createEndpoint bool
)

var (
	availableTransports = []string{"Restful", "WebSockets"}
	defaultTransport    = "Restful"
)

var ResourceCmd = &cobra.Command{
	Use:   "resource <name> [path]",
	Short: "Create a resource with pre-defined components",
	Args:  cobra.MinimumNArgs(1),
	Run:   createResource,
}

func init() {
	ResourceCmd.Flags().StringVarP(&transport, "transport", "t", "", "Available transports are ('Restful', 'WebSockets')")
}

func createResource(cmd *cobra.Command, args []string) {
	resourceName = args[0]

	defaultPath := filepath.Join(".", "app", "modules", resourceName)

	if len(args) > 1 {
		resourcePath = args[1]
	} else {
		resourcePath = defaultPath
	}

	promptForValues()

	generateResourceFromTemplate()

	utils.LogSuccess("Resource created successfully")
}
