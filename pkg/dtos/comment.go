package dtos

// CommentCreation DTO and JSON binding
type CommentCreation struct {
	AppURLID  string `validate:"required,uuid4"`
	PageID    string `validate:"required,uuid4"`
	UserID    int    `validate:"number"`
	Format    string `json:"format" validate:"required" binding:"required"`
	Username  string `json:"username" validate:"required" binding:"required"`
	Body      string `json:"body" validate:"required,min=8,max=1024" binding:"required"`
	Anonymous bool   `json:"anon"`
}
