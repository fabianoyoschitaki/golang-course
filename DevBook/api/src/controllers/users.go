package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// CreateUser creates a new user
func CreateUser(rw http.ResponseWriter, r *http.Request) {
	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		// 422 - unprocessable entity
		responses.Error(rw, http.StatusUnprocessableEntity, error)
		return
	}

	var user models.User
	if error = json.Unmarshal(requestBody, &user); error != nil {
		// 400 - bad request
		responses.Error(rw, http.StatusBadRequest, error)
		return
	}

	// open database connection
	db, error := database.Connect()
	if error != nil {
		// 500 - internal server error
		responses.Error(rw, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	// created repository with database dependency
	repository := repositories.NewUsersRepository(db)
	userID, error := repository.Create(user)
	if error != nil {
		// 500 - internal server error
		responses.Error(rw, http.StatusInternalServerError, error)
		return
	}

	// returning created user
	user.ID = userID
	responses.JSON(rw, http.StatusCreated, user)
}

// FetchUsers fetches all user
func FetchUsers(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Creating user"))
}

// FetchUser fetches a user
func FetchUser(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Creating user"))
}

// UpdateUser updates a user
func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Creating user"))
}

// DeleteUser deletes a new user
func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Creating user"))
}
