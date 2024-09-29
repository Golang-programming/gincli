package service

import "fmt"

func Create() string {
    return fmt.Sprintf("Create operation successfully performed")
}
func GetAll() string {
    return fmt.Sprintf("Fetched all records successfully")
}
func GetById(id string) string {
    return fmt.Sprintf("Fetched record with ID: %s", id)
}
func Update() string {
    return fmt.Sprintf("Update operation successfully performed")
}
func Delete() string {
    return fmt.Sprintf("Delete operation successfully performed")
}
