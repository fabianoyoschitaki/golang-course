package models

// AuthenticationData contains token and id
type AuthenticationData struct {
	Token string `json:"token"`
	ID    string `json:"id"`
}
