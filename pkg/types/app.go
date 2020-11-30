package types

import "time"

// App represents the app instance in database
type App struct {
	ID        int       `json:",omitempty" db:"id"`
	OwnerID   int       `json:",omitempty" db:"owner_id"`
	URLID     string    `json:"UrlId,omitempty" db:"url_id"`
	Name      string    `json:"Name,omitempty" db:"name"`
	Pages     []int     `json:"Pages,omitempty" db:"pages"`
	URL       string    `json:"Url,omitempty" db:"url"`
	CreatedAt time.Time `json:"CreatedAt,omitempty" db:"created_at"`
}
