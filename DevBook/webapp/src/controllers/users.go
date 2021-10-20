package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/requests"
	"webapp/src/responses"

	"github.com/gorilla/mux"
)

// CreateUser calls backend API to create a new user
func CreateUser(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm() // get request body

	// get the form values and convert to a JSON since it's what our API expects
	name := r.FormValue("name")
	nick := r.FormValue("nick")
	email := r.FormValue("email")
	password := r.FormValue("password")
	userToCreate, error := json.Marshal(map[string]string{
		"name":     name,
		"nick":     nick,
		"email":    email,
		"password": password,
	})
	// if we could not marshal the fields to a map, then http 400
	if error != nil {
		responses.JSON(rw, http.StatusBadRequest, responses.APIError{Error: error.Error()})
		return
	}

	// make the request to our backend API
	// if http == 400 or 500, error is nil, because the request was successful! this error means we could not make the request complete
	APIUrl := fmt.Sprintf("%s/users", config.APIURL)
	response, error := http.Post(APIUrl, "application/json", bytes.NewBuffer(userToCreate))
	if error != nil {
		// we cannot use response.statusCode because if error != nil, response doesn't have it! (nil)
		responses.JSON(rw, http.StatusInternalServerError, responses.APIError{Error: error.Error()})
		return
	}
	defer response.Body.Close() // needs to be closed just like database, even if it's empty!

	// if our backend API returned error, then we need to forward it to FE
	if response.StatusCode >= 400 {
		responses.HandleHttpResponseErrors(rw, response)
		return
	}

	// we don't need to return the data from the API to the front-end
	responses.JSON(rw, response.StatusCode, nil)
}

// UnfollowUser calls API to unfollow a user
func UnfollowUser(rw http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userToUnfollowID, error := strconv.ParseUint(parameters["userId"], 10, 64)
	if error != nil {
		responses.JSON(rw, http.StatusBadRequest, responses.APIError{Error: error.Error()})
		return
	}

	unfollowUserAPIUrl := fmt.Sprintf("%s/users/%d/unfollow", config.APIURL, userToUnfollowID)
	APIResponse, error := requests.MakeRequestWithAuthenticationData(r, http.MethodPost, unfollowUserAPIUrl, nil)
	if error != nil {
		responses.JSON(rw, http.StatusInternalServerError, responses.APIError{Error: error.Error()})
		return
	}
	defer APIResponse.Body.Close()

	if APIResponse.StatusCode >= 400 {
		responses.HandleHttpResponseErrors(rw, APIResponse)
		return
	}

	responses.JSON(rw, APIResponse.StatusCode, nil)
}

// FollowUser calls API to unfollow a user
func FollowUser(rw http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	userToFollowID, error := strconv.ParseUint(parameters["userId"], 10, 64)
	if error != nil {
		responses.JSON(rw, http.StatusBadRequest, responses.APIError{Error: error.Error()})
		return
	}

	followUserAPIUrl := fmt.Sprintf("%s/users/%d/follow", config.APIURL, userToFollowID)
	responseAPI, error := requests.MakeRequestWithAuthenticationData(r, http.MethodPost, followUserAPIUrl, nil)
	if error != nil {
		responses.JSON(rw, http.StatusInternalServerError, responses.APIError{Error: error.Error()})
		return
	}
	defer responseAPI.Body.Close()

	if responseAPI.StatusCode >= 400 {
		responses.HandleHttpResponseErrors(rw, responseAPI)
		return
	}

	responses.JSON(rw, responseAPI.StatusCode, nil)
}

// UpdateUser calls API to update a user
func UpdateUser(rw http.ResponseWriter, r *http.Request) {

	// getting user values to be updated
	r.ParseForm()
	newUserValues, error := json.Marshal(map[string]string{
		"name":  r.FormValue("newName"), // they don't come from the form created at edit-profile.html, but from the JSON from users.js
		"nick":  r.FormValue("newNick"),
		"email": r.FormValue("newEmail"),
	})
	if error != nil {
		responses.JSON(rw, http.StatusBadRequest, responses.APIError{Error: error.Error()})
		return
	}

	// getting user ID to be updated
	cookie, _ := cookies.ReadCookie(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	// creating url and request to update user
	updateUserAPIUrl := fmt.Sprintf("%s/users/%d", config.APIURL, userID)

	// making request
	log.Println(string(newUserValues))
	responseAPI, error := requests.MakeRequestWithAuthenticationData(r, http.MethodPut, updateUserAPIUrl, bytes.NewBuffer(newUserValues))
	if error != nil {
		responses.JSON(rw, http.StatusInternalServerError, responses.APIError{Error: error.Error()})
		return
	}
	defer responseAPI.Body.Close()

	if responseAPI.StatusCode >= 400 {
		responses.HandleHttpResponseErrors(rw, responseAPI)
		return
	}

	responses.JSON(rw, responseAPI.StatusCode, nil)
}

// UpdatePassword calls API to update a user's password
func UpdatePassword(rw http.ResponseWriter, r *http.Request) {

	// getting user password values to be updated
	r.ParseForm()
	newPasswordValues, error := json.Marshal(map[string]string{
		"newPassword":     r.FormValue("newPassword"),
		"currentPassword": r.FormValue("currentPassword"),
	})
	if error != nil {
		responses.JSON(rw, http.StatusBadRequest, responses.APIError{Error: error.Error()})
		return
	}

	// getting user ID to be updated
	cookie, _ := cookies.ReadCookie(r)
	userID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	// creating url and request to update password
	updateUserAPIUrl := fmt.Sprintf("%s/users/%d/password-reset", config.APIURL, userID)

	// making request
	responseAPI, error := requests.MakeRequestWithAuthenticationData(r, http.MethodPost, updateUserAPIUrl, bytes.NewBuffer(newPasswordValues))
	if error != nil {
		responses.JSON(rw, http.StatusInternalServerError, responses.APIError{Error: error.Error()})
		return
	}
	defer responseAPI.Body.Close()

	if responseAPI.StatusCode >= 400 {
		responses.HandleHttpResponseErrors(rw, responseAPI)
		return
	}

	responses.JSON(rw, responseAPI.StatusCode, nil)
}

// DeleteUser deletes the user that is logged in
func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.ReadCookie(r)
	loggedUserID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	deleteUserAPIUrl := fmt.Sprintf("%s/users/%d", config.APIURL, loggedUserID)
	responseAPI, error := requests.MakeRequestWithAuthenticationData(r, http.MethodDelete, deleteUserAPIUrl, nil)
	if error != nil {
		responses.JSON(rw, http.StatusInternalServerError, responses.APIError{Error: error.Error()})
		return
	}
	defer responseAPI.Body.Close()

	if responseAPI.StatusCode >= 400 {
		responses.HandleHttpResponseErrors(rw, responseAPI)
		return
	}

	responses.JSON(rw, responseAPI.StatusCode, nil)
}
