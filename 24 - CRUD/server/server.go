package server

import (
	"basic-crud/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// lowercase because we'll use it only inside this file
type user struct {
	ID    uint32 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// FetchUsers fetches all users from database
func FetchUsers(rw http.ResponseWriter, r *http.Request) {
	db, error := database.Connect()
	if error != nil {
		rw.Write([]byte("Error when connecting to database"))
		return
	}
	defer db.Close()

	// SELECT * FROM USERS
	rows, error := db.Query("select * from users")
	if error != nil {
		rw.Write([]byte("Error when fetching users"))
		return
	}
	defer rows.Close()

	var users []user
	for rows.Next() {
		var user user

		// ID  NAME EMAIL
		if error := rows.Scan(&user.ID, &user.Name, &user.Email); error != nil {
			rw.Write([]byte("Error when scanning user"))
			return
		}

		users = append(users, user)
	}
	rw.WriteHeader(http.StatusOK)
	if error := json.NewEncoder(rw).Encode(users); error != nil {
		rw.Write([]byte("Error when converting users to JSON"))
		return
	}

}

// FetchUsers fetches all users from database
func FetchUser(rw http.ResponseWriter, r *http.Request) {

}

// CreateUser creates a new user
func CreateUser(rw http.ResponseWriter, r *http.Request) {

	// get request body (the user to be inserted) {"name" : "John", "email": "john@doe.com"}
	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		rw.Write([]byte("Failure to read request body"))
		return // stop execution
	}

	// get request body and unmarshal it to struct
	var user user
	if error = json.Unmarshal(requestBody, &user); error != nil {
		rw.Write([]byte("Error when converting user to struct"))
		return
	}

	// just printing the user to be inserted
	fmt.Println(user)
	// POST localhost:5000/users
	// {"name" : "John", "email": "john@doe.com"}

	db, error := database.Connect()
	if error != nil {
		rw.Write([]byte("Error when connecting to database"))
		fmt.Println(error)
		return // we dont need to test ping since it's being done in the Connnect function
	}
	defer db.Close() // it must be closed

	// INSERT INTO USERS (NAME, EMAIL) VALUES (X, Y)
	// let's use preparared statement to avoid SQL injections
	statement, error := db.Prepare("insert into users (name, email) values (?, ?)")
	if error != nil {
		rw.Write([]byte("Error when creating statement"))
		return
	}
	defer statement.Close() // it must be closed just like db\

	// inserting real data (the values)
	insertion, error := statement.Exec(user.Name, user.Email)
	if error != nil {
		rw.Write([]byte("Error when executing statement"))
		return
	}

	// user was inserted at this point. Let's return the ID inserted
	insertedId, error := insertion.LastInsertId()
	if error != nil {
		rw.Write([]byte("Error when fetching inserted ID"))
		return
	}

	// response and status codes
	// rw.WriteHeader(201) // CREATED
	rw.WriteHeader(http.StatusCreated) // same thing as above, just a constant
	rw.Write([]byte(fmt.Sprintf("User added successfully with ID: %d", insertedId)))

}
