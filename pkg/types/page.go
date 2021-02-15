package types

// Page represents the system page
type Page struct {
	ID       int    `json:"ID,-,omitempty"`
	PageID   string `json:"PageId" db:"page_id"`
	AppURLID string `json:",omitempty,appUrlId"`
	Title    string `json:"Title"`
	Slug     string `json:"Slug"`
	URL      string `json:"URL"`
}

// IsEmpty checks if the structure is empty
func (p Page) IsEmpty() bool {
	return len(p.PageID) <= 0
}
