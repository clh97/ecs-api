package dtos

// Login DTO and JSON binding
type Login struct {
	Email    string `json:"email" validate:"required,email" binding:"required"`
	Password string `json:"password" validate:"required" binding:"required"`
}

// AccountCreation DTO and JSON binding
type AccountCreation struct {
	Email    string `json:"email" validate:"required,email" binding:"required"`
	Password string `json:"password" validate:"required,min=8,max=128" binding:"required"`
	Username string `json:"username" validate:"required" binding:"required"`
}
