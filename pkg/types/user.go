package types

// User represents the system user
type User struct {
	ID       int
	Email    string `json:"Email"`
	Password string `json:"Password,-"`
	Username string `json:"Username"`
}
