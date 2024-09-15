package resource

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	resourceName string
	transport    string
	skipPrompts  bool
)
var availableTransports = []string{"restful", "graphql", "webSockets"}

var (
	defaultTransport = "restful"
)

var ResourceCmd = &cobra.Command{
	Use:     "resource",
	Short:   "Create a resource with pre-defined components",
	Run:     createResource,
	Aliases: []string{"res", "create"},
}

func init() {
	ResourceCmd.Flags().StringVar(&transport, "transport", "", "Available transports are ('restful', 'graphql', 'webSockets')")
	ResourceCmd.Flags().BoolVarP(&skipPrompts, "yes", "y", false, "Skip all prompts and use default values")
}

func createResource(cmd *cobra.Command, args []string) {
	if skipPrompts {
		setDefaultValues()
	} else {
		promptForValues()
	}

	createResourceFromTemplate()

	// runGoModTidy(projectDir)

	fmt.Println(color.New(color.FgGreen).Sprint("Resource created successfully"))
}

func setDefaultValues() {
	if transport == "" {
		transport = defaultTransport
	}
}
