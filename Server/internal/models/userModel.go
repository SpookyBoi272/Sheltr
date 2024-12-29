package models

type User struct {
	ID           int    `json:"user_id,omitempty"`
	Email        string `json:"email,omitempty"`
	Password     string `json:"password,omitempty"`
	FirstName    string `json:"first-name,omitempty"`
	LastName     string `json:"last-name,omitempty"`
	PasswordHash string `json:"password-hash,omitempty"`
}
