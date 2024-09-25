package dtos

type Create{{.CapitalizeResourceName}}Input struct {
	Name string `json:"name" binding:"required"`
	// Add more fields as necessary
}

type Update{{.CapitalizeResourceName}}Input struct {
	Name string `json:"name" binding:"omitempty"`
	// Add more fields as necessary
}
