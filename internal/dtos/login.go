package dtos

// Login DTO and JSON binding
type Login struct {
	Email    string `json:"email" validate:"required,email" binding:"required"`
	Password string `json:"password" validate:"required" binding:"required"`
}
