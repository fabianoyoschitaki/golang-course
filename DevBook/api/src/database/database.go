package database

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // implicit import for Driver
)

// Connect opens database connection and returns it
func Connect() (*sql.DB, error) {
	db, error := sql.Open("mysql", config.DatabaseConnectionString)
	if error != nil {
		return nil, error
	}
	// no defer db.close() because we want to return the connection opened to whoever called

	if error = db.Ping(); error != nil {
		db.Close() // we want to close, because we had a problem
		return nil, error
	}

	// everything is alright now
	return db, nil
}
