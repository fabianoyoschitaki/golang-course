package models

import "time"

// User represents a user
type User struct {
	ID        uint64    `json:"id,omitempty"` // if id is blank, it won't be considered when parsing to JSON instead of 0
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}
