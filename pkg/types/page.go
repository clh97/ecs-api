package types

// Page represents the system page
type Page struct {
	ID     int    `json:"id,-"`
	PageID string `json:"pageId"`
	AppID  string `json:"appId"`
	Title  string `json:"title"`
	URL    string `json:"url"`
}

// IsEmpty checks if the structure is empty
func (p Page) IsEmpty() bool {
	return p.ID <= 0
}
