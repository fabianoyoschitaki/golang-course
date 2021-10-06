package models

import (
	"errors"
	"strings"
	"time"
)

// User represents a user
type User struct {
	ID        uint64    `json:"id,omitempty"` // if id is blank, it won't be considered when parsing to JSON instead of 0
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

// Prepare validates and formats user
func (u *User) Prepare() error {
	if error := u.validate(); error != nil {
		return error
	}

	u.format()
	return nil
}

func (u *User) validate() error {
	if u.Name == "" {
		return errors.New("User name can't be blank")
	}
	if u.Nick == "" {
		return errors.New("User nick can't be blank")
	}
	if u.Email == "" {
		return errors.New("User email can't be blank")
	}
	if u.Password == "" {
		return errors.New("User password can't be blank	")
	}
	return nil
}

func (u *User) format() {
	u.Name = strings.TrimSpace(u.Name)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)
	// not email since space could be on purpose
}
