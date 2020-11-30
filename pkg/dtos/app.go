package dtos

// AppCreation DTO and JSON binding
type AppCreation struct {
	Name string `json:"name" validate:"required" binding:"required"`
	URL  string `json:"url" validate:"required,url" binding:"required"`
}
