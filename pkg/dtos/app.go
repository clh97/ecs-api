package dtos

// AppCreation DTO and JSON binding
type AppCreation struct {
	Name string `json:"name" validate:"required" binding:"required"`
	URL  string `json:"url" validate:"required,url" binding:"required"`
}
// AppDelete DTO and JSON binding
type AppDelete struct {
	URLID string `json:"urlId" validate:"required,uuid4" binding:"required"`
}
