package repositories

import (
	"api/src/models"
	"database/sql"
)

// it's lowercase because we won't export it
type users struct {
	db *sql.DB // will receive a database from controllers
}

// NewUsersRepository creates users repository
func NewUsersRepository(db *sql.DB) *users {
	// inside this struct we'll have the database operations, insert, update etc.
	// #IMPORTANT: controller only opens connection, repository makes connection with tables
	return &users{db}
}

// Create inserts a new user in the database
func (repository users) Create(user models.User) (uint64, error) {
	statement, error := repository.db.Prepare("insert into users (name, nick, email, password) values (?, ?, ?, ?)")
	if error != nil {
		return 0, error
	}
	defer statement.Close()

	result, error := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if error != nil {
		return 0, error
	}

	// at this point the user is already inserted into database
	// get ID
	lastInsertedID, error := result.LastInsertId()
	if error != nil {
		return 0, nil
	}
	return uint64(lastInsertedID), nil
}
