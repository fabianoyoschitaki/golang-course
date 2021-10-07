package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type Users struct {
	db *sql.DB // will receive a database from controllers
}

// NewUsersRepository creates Users repository
func NewUsersRepository(db *sql.DB) *Users {
	// inside this struct we'll have the database operations, insert, update etc.
	// #IMPORTANT: controller only opens connection, repository makes connection with tables
	return &Users{db}
}

// Search returns all users which has name or nick
func (repo Users) Search(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick) // %nameOrNick%, we need to scape %

	// query where name or nick like
	rows, error := repo.db.Query("select id, name, nick, email, created_at from users where name like ? or nick like ?",
		nameOrNick, nameOrNick)
	if error != nil {
		return nil, error
	}
	defer rows.Close()

	// response
	var usersFound []models.User
	for rows.Next() {
		var currentUser models.User
		if error = rows.Scan(
			&currentUser.ID,
			&currentUser.Name,
			&currentUser.Nick,
			&currentUser.Email,
			&currentUser.CreatedAt,
		); error != nil {
			return nil, error
		}
		usersFound = append(usersFound, currentUser)
	}
	return usersFound, nil
}

// FetchUserByID fetches a single user by its ID
func (repo Users) FetchUserByID(userID uint64) (models.User, error) {
	// query where name or nick like
	rows, error := repo.db.Query(
		"select id, name, nick, email, created_at from users where id = ?", userID,
	)

	if error != nil {
		return models.User{}, error
	}
	defer rows.Close()

	// response
	var userResponse models.User
	if rows.Next() {
		if error = rows.Scan(
			&userResponse.ID,
			&userResponse.Name,
			&userResponse.Nick,
			&userResponse.Email,
			&userResponse.CreatedAt,
		); error != nil {
			return models.User{}, error
		}
	}
	return userResponse, nil
}

// Update updates a user in the database
func (repo Users) Update(ID uint64, userToUpdate models.User) error {
	statement, error := repo.db.Prepare(("update users set name = ?, nick = ?, email = ? where id = ?"))
	if error != nil {
		return error
	}
	defer statement.Close()

	_, error = statement.Exec(userToUpdate.Name, userToUpdate.Nick, userToUpdate.Email, ID)
	if error != nil {
		return nil
	}
	// everything is alright
	return nil
}

// Delete deletes a user from the database
func (repo Users) Delete(ID uint64) error {
	statement, error := repo.db.Prepare("delete from users where id = ?")
	if error != nil {
		return error
	}
	defer statement.Close()

	_, error = statement.Exec(ID)
	if error != nil {
		return error
	}
	return nil
}

// Create inserts a new user in the database
func (repo Users) Create(user models.User) (uint64, error) {
	statement, error := repo.db.Prepare("insert into users (name, nick, email, password) values (?, ?, ?, ?)")
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
