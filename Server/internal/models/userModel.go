package models

type User struct {
	ID           int    `json:"id,omitempty"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	FirstName    string `json:"first-name"`
	LastName     string `json:"last-name"`
	PasswordHash string `json:"password-hash,omitempty"`
}
