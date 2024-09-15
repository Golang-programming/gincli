package services

import (
    "fmt"
    "{{.AppName}}/app/modules/{{.ResourceName}}/dtos"
)

func CreateCapitalizeResourceName(input dtos.CreateCapitalizeResourceNameInput) string {
    return fmt.Sprintf("Create {{.ResourceName}} successfully")
}

func GetAllCapitalizeResourceNames() string {
    return fmt.Sprintf("Fetched all  {{.ResourceName}} successfully")
}

func GetCapitalizeResourceNameById(ID string) string {
    return fmt.Sprintf("Fetched  {{.ResourceName}} with ID: %s", ID)
}

func UpdateCapitalizeResourceName(ID string,input dtos.UpdateCapitalizeResourceNameInput) string {
    return fmt.Sprintf("Update  {{.ResourceName}} successfully")
}

func DeleteCapitalizeResourceName(ID string, ) string {
    return fmt.Sprintf("Delete  {{.ResourceName}} successfully")
}