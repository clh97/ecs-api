package dtos

// PageCreation DTO and Binding
type PageCreation struct {
	AppURLID string
	PageID   int
	Title    string
	URL      string
}

// PageGet DTO and Binding
type PageGet struct {
	AppURLID string
	PageID   int
}
