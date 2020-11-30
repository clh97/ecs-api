package types

import "time"

// App represents the app instance in database
type App struct {
	ID        int       `json:",omitempty" db:"id"`
	OwnerID   int       `json:",omitempty" db:"owner_id"`
	URLID     string    `json:"url_id,omitempty" db:"url_id"`
	Name      string    `json:"name,omitempty" db:"name"`
	Pages     []int     `json:",omitempty" db:"pages"`
	URL       string    `json:"url,omitempty" db:"url"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
}
