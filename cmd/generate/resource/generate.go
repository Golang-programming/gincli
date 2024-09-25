// cmd/generate/resource/generate.go
package resource

/*
func createResource(cmd *cobra.Command, args []string) {
	resourceName = args[0]

	promptForValues()

	projectDir := filepath.Join(".", "app", "modules", resourceName)
	utils.CreateDirectories([]string{
		filepath.Join(projectDir, "controllers"),
		filepath.Join(projectDir, "services"),
		filepath.Join(projectDir, "dtos"),
		filepath.Join(projectDir, "entities"),
	})

	generateController(projectDir)
	generateService(projectDir)
	generateDTOs(projectDir)
	generateEntities(projectDir)

	if createEndpoints {
		generateRoutes(projectDir)
	}

	fmt.Println(color.New(color.FgGreen).Sprint("Resource generated successfully"))
}

func generateController(projectDir string) {
	templatePath := "templates/generate/resource/controller.go.tpl"
	outputPath := filepath.Join(projectDir, "controllers", "controller.go")
	config := map[string]string{
		"Module":                 utils.DetectModuleName(),
		"ResourceName":           resourceName,
		"CapitalizeResourceName": utils.Capitalize(resourceName),
	}
	utils.GenerateFileFromTemplate(templatePath, outputPath, config)
}

func generateService(projectDir string) {
	templatePath := "templates/generate/resource/service.go.tpl"
	outputPath := filepath.Join(projectDir, "services", "service.go")
	config := map[string]string{
		"Module":                 utils.DetectModuleName(),
		"ResourceName":           resourceName,
		"CapitalizeResourceName": utils.Capitalize(resourceName),
	}
	utils.GenerateFileFromTemplate(templatePath, outputPath, config)
}

func generateDTOs(projectDir string) {
	templatePath := "templates/generate/resource/dto.go.tpl"
	outputPath := filepath.Join(projectDir, "dtos", "dto.go")
	config := map[string]string{
		"ResourceName":           resourceName,
		"CapitalizeResourceName": utils.Capitalize(resourceName),
	}
	utils.GenerateFileFromTemplate(templatePath, outputPath, config)
}

func generateEntities(projectDir string) {
	templatePath := "templates/generate/resource/entity.go.tpl"
	outputPath := filepath.Join(projectDir, "entities", "entity.go")
	config := map[string]string{
		"ResourceName":           resourceName,
		"CapitalizeResourceName": utils.Capitalize(resourceName),
	}
	utils.GenerateFileFromTemplate(templatePath, outputPath, config)
}

func generateRoutes(projectDir string) {
	templatePath := "templates/generate/resource/routes.go.tpl"
	outputPath := filepath.Join(projectDir, "routes.go")
	config := map[string]string{
		"Module":                 utils.DetectModuleName(),
		"ResourceName":           resourceName,
		"CapitalizeResourceName": utils.Capitalize(resourceName),
	}
	utils.GenerateFileFromTemplate(templatePath, outputPath, config)
}
*/
