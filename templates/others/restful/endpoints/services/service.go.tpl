package services

import (
    "fmt"
    "{{.Module}}/app/modules/{{.ResourceName}}/dtos"
)

func Create{{.CapitalizeResourceName}}(input dtos.Create{{.CapitalizeResourceName}}Input) string {
    return fmt.Sprintf("Create {{.ResourceName}} successfully")
}

func GetAll{{.CapitalizeResourceName}}s() string {
    return fmt.Sprintf("Fetched all  {{.ResourceName}} successfully")
}

func Get{{.CapitalizeResourceName}}ById(ID string) string {
    return fmt.Sprintf("Fetched  {{.ResourceName}} with ID: %s", ID)
}

func Update{{.CapitalizeResourceName}}(ID string,input dtos.Update{{.CapitalizeResourceName}}Input) string {
    return fmt.Sprintf("Update  {{.ResourceName}} successfully")
}

func Delete{{.CapitalizeResourceName}}(ID string, ) string {
    return fmt.Sprintf("Delete  {{.ResourceName}} successfully")
}