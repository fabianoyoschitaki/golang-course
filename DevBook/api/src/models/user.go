package models

import (
	"api/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
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
func (u *User) Prepare(step string) error {
	if err := u.validate(step); err != nil {
		return err
	}

	if err := u.format(step); err != nil {
		return err
	}
	return nil
}

func (u *User) validate(step string) error {
	if u.Name == "" {
		return errors.New("User name can't be blank")
	}
	if u.Nick == "" {
		return errors.New("User nick can't be blank")
	}
	if u.Email == "" {
		return errors.New("User email can't be blank")
	}
	if error := checkmail.ValidateFormat(u.Email); error != nil {
		// could return same error
		return errors.New("User email is invalid")
	}

	if step == "signup" && u.Password == "" {
		return errors.New("User password can't be blank	")
	}
	return nil
}

func (u *User) format(step string) error {
	u.Name = strings.TrimSpace(u.Name)
	u.Nick = strings.TrimSpace(u.Nick)
	u.Email = strings.TrimSpace(u.Email)
	// don't do it for the password since space could be on purpose

	// if we're signing up, then we should hash the informed password with BCrypt
	if step == "signup" {
		hashedPassword, err := security.Hash(u.Password)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}
	return nil
}
