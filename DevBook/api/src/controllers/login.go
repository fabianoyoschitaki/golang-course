package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Login is responsible to authenticate a user within API
func Login(rw http.ResponseWriter, r *http.Request) {
	// receive request with email + password, fetch from email and check if hashed password matches
	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		responses.Error(rw, http.StatusUnprocessableEntity, error)
		return
	}

	// the user from request to struct
	var loginUser models.User
	if error = json.Unmarshal(requestBody, &loginUser); error != nil {
		responses.Error(rw, http.StatusBadRequest, error)
		return
	}

	// fetching user from database
	db, error := database.Connect()
	if error != nil {
		responses.Error(rw, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	// fetching ID and password for user
	userRepository := repositories.NewUsersRepository(db)
	userFromDatabase, error := userRepository.FetchUserByEmail(loginUser.Email)
	if error != nil {
		responses.Error(rw, http.StatusInternalServerError, error)
		return
	}

	// let's compare user from database with provided password
	if error = security.CheckPassword(loginUser.Password, userFromDatabase.Password); error != nil {
		responses.Error(rw, http.StatusUnauthorized, error)
		return
	}

	// password matches. Generate JWT
	token, error := authentication.CreateToken(userFromDatabase.ID)
	if error != nil {
		responses.Error(rw, http.StatusInternalServerError, error)
		return
	}
	rw.Write([]byte(token))
}
