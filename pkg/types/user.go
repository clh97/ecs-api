package types

// User represents the system user
type User struct {
	ID        int
	Email     string `json:"email"`
	Password  string `json:"password,-"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
