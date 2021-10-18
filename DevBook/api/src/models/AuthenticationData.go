package models

// AuthenticationData returns the user ID and JWT
type AuthenticationData struct {
	Token string `json:"token"`
	ID    string `json:"id"`
}
