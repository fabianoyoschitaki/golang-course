package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // this is manual. MYSQL connection driver
)

// Connect opens a database connection with a MySQL database
func Connect() (*sql.DB, error) {
	connectionURL := "golang:password123@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, error := sql.Open("mysql", connectionURL)
	if error != nil {
		return nil, error // mutually exclusive
	}
	// we don't defer the db, only who calls
	if error = db.Ping(); error != nil {
		return nil, error
	}

	// if everything goes alright
	return db, nil
}
