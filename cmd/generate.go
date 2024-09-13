package cmd

import (
	"fmt"

	"github.com/golang-programming/gincli/utils"
	"github.com/spf13/cobra"
)

var componentType, name string

var generateCmd = &cobra.Command{
    Use:   "generate",
    Short: "Generate a new component (controller, service, model, etc.)",
    Run: func(cmd *cobra.Command, args []string) {
        if len(args) < 2 {
            fmt.Println("Please provide the component type and name.")
            return
        }
        componentType = args[0]
        name = args[1]
        generateComponent(componentType, name)
    },
}

func init() {
    rootCmd.AddCommand(generateCmd)
}

func generateComponent(compType, name string) {
    fmt.Printf("Generating %s: %s\n", compType, name)
    templatePath := fmt.Sprintf("./templates/%s.tpl", compType)
    outputPath := fmt.Sprintf("./%s/%s.go", compType, name)

    utils.GenerateFileFromTemplate(templatePath, outputPath, map[string]string{
        "Name": name,
    })
}
