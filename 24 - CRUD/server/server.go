package server

import (
	"basic-crud/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

		// ID  NAME  EMAIL
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

// FetchUser fetches a user by its ID
func FetchUser(rw http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	// getting parameter in decimal base with 32 bits
	ID, error := strconv.ParseUint(parameters["id"], 10, 32)
	if error != nil {
		rw.Write([]byte("Error when converting parameter to integer"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		rw.Write([]byte("Error when connecting to database"))
		return
	}
	defer db.Close()

	row, error := db.Query("select * from users where id = ?", ID)
	if error != nil {
		rw.Write([]byte("Error when fetching user"))
		return
	}

	var user user
	if row.Next() {
		if error := row.Scan(&user.ID, &user.Name, &user.Email); error != nil {
			rw.Write([]byte("Error when scanning user"))
			return
		}
	}

	rw.WriteHeader(http.StatusOK)
	if error := json.NewEncoder(rw).Encode(user); error != nil {
		rw.Write([]byte("Error when converting user to JSON"))
		return
	}
}

// UpdateUser updates a user
func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	ID, error := strconv.ParseUint(parameters["id"], 10, 32)
	if error != nil {
		rw.Write([]byte("Error when converting parameter to integer"))
		return
	}

	// read request body, so that only in the end we open database connection
	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		rw.Write([]byte("Error when reading request body"))
		return
	}

	var user user
	if error := json.Unmarshal(requestBody, &user); error != nil {
		rw.Write([]byte("Error when converting user to struct"))
		return
	}

	// at this point we have the ID and the user request body
	db, error := database.Connect()
	if error != nil {
		rw.Write([]byte("Error when connecting to database"))
		return
	}
	defer db.Close()

	statement, error := db.Prepare("update users set name = ?, email = ? where id = ?")
	if error != nil {
		rw.Write([]byte("Error when creating statement"))
		return
	}
	defer statement.Close()

	if _, error := statement.Exec(user.Name, user.Email, ID); error != nil {
		rw.Write([]byte("Error when updating user"))
		return
	}
	// if everything went well we just return http 204
	rw.WriteHeader(http.StatusNoContent)
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

// DeleteUser deletes a user
func DeleteUser(rw http.ResponseWriter, r *http.Request) {

	parameters := mux.Vars(r)
	ID, error := strconv.ParseUint(parameters["id"], 10, 32)
	if error != nil {
		rw.Write([]byte("Error when converting parameter to integer"))
		return
	}

	db, error := database.Connect()
	if error != nil {
		rw.Write([]byte("Error when connecting to database"))
		return
	}
	defer db.Close()

	statement, error := db.Prepare("delete from users where id = ?")
	if error != nil {
		rw.Write([]byte("Error when creating statement"))
		return
	}
	defer statement.Close()

	if _, error := statement.Exec(ID); error != nil {
		rw.Write([]byte("Error when deleting user"))
		return
	}
	// if everything went well we just return http 204
	rw.WriteHeader(http.StatusNoContent)
}
