package dtos

// PageCreation DTO and Binding
type PageCreation struct {
	AppURLID string `json:"appUrlId" validate:"required,uuid4" binding:"required"`
	Title    string `json:"title" validate:"required" binding:"required"`
	URL      string `json:"url" validate:"required" binding:"required"`
}

// PageGet DTO and Binding
type PageGet struct {
	AppURLID string
	PageID   string
}
