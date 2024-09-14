package dtos

type Create{{Capitalize .ResourceName}}Input struct {
	Name         string            `json:"name" binding:"required"`
}


type Update{{Capitalize .ResourceName}}Input struct {
	Name         string            `json:"name" binding:"omitempty"`
}
