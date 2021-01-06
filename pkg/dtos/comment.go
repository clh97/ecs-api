package dtos

// CommentCreation DTO and JSON binding
type CommentCreation struct {
	Username  string `json:"username" validate:"required" binding:"required"`
	Body      string `json:"body" validate:"required,min=8,max=1024" binding:"required"`
	Email     string `json:"email" validate:"email"`
	Anonymous bool   `json:"anon" binding:"required"`
}
