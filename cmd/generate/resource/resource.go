package resource

import (
	"fmt"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	resourceName string
	resourcePath string
	transport    string
)

var (
	availableTransports = []string{"Restful", "WebSockets"}
	defaultTransport    = "Restful"
)

var ResourceCmd = &cobra.Command{
	Use:   "resource <name> [path]",
	Short: "Create a resource with pre-defined components",
	Run:   createResource,
	Args:  cobra.MinimumNArgs(1),
}

func init() {
	ResourceCmd.Flags().StringVar(&transport, "transport", "", "Available transports are ('Restful', 'WebSockets')")
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

	createResourceFromTemplate()

	fmt.Println(color.New(color.FgGreen).Sprint("Resource created successfully"))
}
