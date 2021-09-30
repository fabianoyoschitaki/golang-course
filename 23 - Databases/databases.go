package main

import (
	"database/sql" // this will search for mysql
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// #IMPORTANT go get github.com/go-sql-driver/mysql

	// open connection with MYSQL:
	connectionUrl := "golang:password123@/devbook?charset=utf8&parseTime=True&loc=Local"

	// sql: unknown driver "mysql" (forgotten import?) # implicit import
	db, error := sql.Open("mysql", connectionUrl)
	if error != nil {
		log.Fatal(error)
	}
	defer db.Close()

	// how to check if we're connected? (reusing error variable, that's why we dont need := )
	if error = db.Ping(); error != nil {
		fmt.Println("Ping error...")
		log.Fatal(error)
	}
	fmt.Println("Connection is open!")

	// execute SQL
	rows, error := db.Query("select * from users")
	if error != nil {
		log.Fatal(error)
	}
	// we always close the rows (cursor) as well as the database
	defer rows.Close()

	// show lines
	fmt.Println(rows)

}
