package dtos

type Create{{.CapitalizeResourceName}}Input struct {
	Name         string            `json:"name" binding:"required"`
}


type Update{{.CapitalizeResourceName}}Input struct {
	Name         string            `json:"name" binding:"omitempty"`
}
