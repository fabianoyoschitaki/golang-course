package controllers

import (
	"api/src/authentication"
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// CreateUser creates a new user
func CreateUser(rw http.ResponseWriter, r *http.Request) {

	// read request body
	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		// 422 - unprocessable entity
		responses.Error(rw, http.StatusUnprocessableEntity, error)
		return
	}

	// unmarshal to struct
	var user models.User
	if error = json.Unmarshal(requestBody, &user); error != nil {
		// 400 - bad request
		responses.Error(rw, http.StatusBadRequest, error)
		return
	}

	// Validate user according to our rules
	if error = user.Prepare("signup"); error != nil {
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

	// create repository with database dependency and save new user
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

	// user query parameter
	nameOrNick := strings.ToLower(r.URL.Query().Get("user"))

	db, error := database.Connect()
	if error != nil {
		responses.Error(rw, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	repository := repositories.NewUsersRepository(db)
	usersFound, error := repository.Search(nameOrNick)
	if error != nil {
		responses.Error(rw, http.StatusInternalServerError, error)
		return
	}

	// write slice of users as json
	responses.JSON(rw, http.StatusOK, usersFound)
}

// FetchUser fetches a user by its ID
func FetchUser(rw http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r) // parameters: {id}

	// getting id path parameter
	userID, error := strconv.ParseUint(parameters["id"], 10, 64)
	if error != nil {
		responses.Error(rw, http.StatusBadRequest, error) // bad request because ID should be number, not string.
		return
	}

	// open connection
	db, error := database.Connect()
	if error != nil {
		responses.Error(rw, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	// create repository
	repository := repositories.NewUsersRepository(db)
	userByID, error := repository.FetchUserByID(userID)
	if error != nil {
		responses.Error(rw, http.StatusInternalServerError, error)
		return
	}

	// write user as json
	responses.JSON(rw, http.StatusOK, userByID)
}

// UpdateUser updates a user
func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r) // parameters: {id}

	// getting id path parameter
	userID, error := strconv.ParseUint(parameters["id"], 10, 64)
	if error != nil {
		responses.Error(rw, http.StatusBadRequest, error) // bad request because ID should be number, not string.
		return
	}

	// validating if logged user is the same being updated
	loggedUserID, error := authentication.ExtractUserId(r)
	if error != nil {
		responses.Error(rw, http.StatusUnauthorized, error)
		return
	}
	// User is authenticated, but it's forbidden (403) due to not being allowed to update another user
	if loggedUserID != userID {
		responses.Error(rw, http.StatusForbidden, errors.New("it's not possible to update a different user"))
		return
	}

	// read request body
	requestBody, error := ioutil.ReadAll(r.Body)
	if error != nil {
		responses.Error(rw, http.StatusUnprocessableEntity, error)
		return
	}

	// request body to struct
	var userToUpdate models.User
	if error = json.Unmarshal(requestBody, &userToUpdate); error != nil {
		responses.Error(rw, http.StatusBadRequest, error)
		return
	}

	// Validate user according to our rules
	if error = userToUpdate.Prepare("update"); error != nil {
		// 400 - bad request
		responses.Error(rw, http.StatusBadRequest, error)
		return
	}

	// open connection
	db, error := database.Connect()
	if error != nil {
		responses.Error(rw, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	// create repository
	repository := repositories.NewUsersRepository(db)
	if error = repository.Update(userID, userToUpdate); error != nil {
		responses.Error(rw, http.StatusInternalServerError, error)
		return
	}

	// return 204
	responses.JSON(rw, http.StatusNoContent, nil)
}

// DeleteUser deletes a new user
func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r) // parameters: {id}

	// getting id path parameter
	userID, error := strconv.ParseUint(parameters["id"], 10, 64)
	if error != nil {
		responses.Error(rw, http.StatusBadRequest, error) // bad request because ID should be number, not string.
		return
	}

	// only own user could delete himself
	loggedUserID, error := authentication.ExtractUserId(r)
	if error != nil {
		responses.Error(rw, http.StatusUnauthorized, error)
		return
	}
	if loggedUserID != userID {
		responses.Error(rw, http.StatusForbidden, errors.New("it's not possible to delete a user that is not you"))
		return
	}

	// open connection
	db, error := database.Connect()
	if error != nil {
		responses.Error(rw, http.StatusInternalServerError, error)
		return
	}
	defer db.Close()

	// create repository
	repository := repositories.NewUsersRepository(db)
	if error = repository.Delete(userID); error != nil {
		responses.Error(rw, http.StatusInternalServerError, error)
		return
	}

	// return 204
	responses.JSON(rw, http.StatusNoContent, nil)
}
