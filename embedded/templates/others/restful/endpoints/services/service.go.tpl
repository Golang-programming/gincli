package services

import (
	"fmt"
	"{{.Module}}/app/modules/{{.ResourceName}}/dtos"
)

func Create{{.CapitalizeResourceName}}(input dtos.Create{{.CapitalizeResourceName}}Input) string {
	// Implement your create logic here
	return fmt.Sprintf("Create {{.ResourceName}} successfully")
}

func GetAll{{.CapitalizeResourceName}}s() string {
	// Implement your get all logic here
	return fmt.Sprintf("Fetched all {{.ResourceName}} successfully")
}

func Get{{.CapitalizeResourceName}}ById(ID string) string {
	// Implement your get by ID logic here
	return fmt.Sprintf("Fetched {{.ResourceName}} with ID: %s", ID)
}

func Update{{.CapitalizeResourceName}}(ID string, input dtos.Update{{.CapitalizeResourceName}}Input) string {
	// Implement your update logic here
	return fmt.Sprintf("Update {{.ResourceName}} successfully")
}

func Delete{{.CapitalizeResourceName}}(ID string) string {
	// Implement your delete logic here
	return fmt.Sprintf("Delete {{.ResourceName}} successfully")
}
