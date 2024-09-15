package dtos

type CreateCapitalizeResourceNameInput struct {
	Name         string            `json:"name" binding:"required"`
}


type UpdateCapitalizeResourceNameInput struct {
	Name         string            `json:"name" binding:"omitempty"`
}
