// cmd/generate/controller/controller.go
package controller

import (
	"github.com/spf13/cobra"
)

var (
	controllerName string
	skipPrompts    bool
)

var ControllerCmd = &cobra.Command{
	Use:     "controller <name>",
	Short:   "Generate a new controller",
	Args:    cobra.ExactArgs(1),
	Run:     createController,
	Aliases: []string{"ctrl"},
}

func init() {
	ControllerCmd.Flags().BoolVarP(&skipPrompts, "yes", "y", false, "Skip prompts and use default values")
}
