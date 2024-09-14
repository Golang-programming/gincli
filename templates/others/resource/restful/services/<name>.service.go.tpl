package services

import (
    "fmt"
    "{{.AppName}}/app/modules/{{.Resource}}/dtos"
)

func Create{{Capitalize .ResourceName}}(input dtos.Create{{Capitalize .ResourceName}}Input) string {
    return fmt.Sprintf("Create {{.ResourceName}} successfully")
}

func GetAll{{Capitalize .ResourceName}}s() string {
    return fmt.Sprintf("Fetched all  {{.ResourceName}} successfully")
}

func Get{{Capitalize .ResourceName}}ById(ID string) string {
    return fmt.Sprintf("Fetched  {{.ResourceName}} with ID: %s", ID)
}

func Update{{Capitalize .ResourceName}}(ID string,input dtos.Update{{Capitalize .ResourceName}}Input) string {
    return fmt.Sprintf("Update  {{.ResourceName}} successfully")
}

func Delete{{Capitalize .ResourceName}}(ID string, ) string {
    return fmt.Sprintf("Delete  {{.ResourceName}} successfully")
}