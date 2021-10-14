package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/models"
	"webapp/src/responses"
)

// AttemptLogin uses email and password to authenticate against our backend
func AttemptLogin(rw http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	login, error := json.Marshal(map[string]string{
		"email":    r.FormValue("email"),
		"password": r.FormValue("password"),
	})
	// if we cannot marshal to a map, we return a HTTP 400 bad request
	if error != nil {
		responses.JSON(rw, http.StatusBadRequest, responses.APIError{Error: error.Error()})
	}

	// getting API url from config
	APIUrl := fmt.Sprintf("%s/login", config.APIURL)
	response, error := http.Post(APIUrl, "application/json", bytes.NewBuffer(login))
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

	// we're returning from backend API /login route:
	// {
	// 	  "token": "eyJhbGc...",
	//    "id": 1
	// }
	// rather then only the JWT (otherwise we would need to replicate code to get the JWT claims)
	var authenticationData models.AuthenticationData
	if error := json.NewDecoder(response.Body).Decode(&authenticationData); error != nil {
		responses.JSON(rw, http.StatusUnprocessableEntity, responses.APIError{Error: error.Error()})
		return
	}

	// #IMPORTANT
	// user is authenticated. now that we have the JWT and the user ID, we need to save it in the user's browser as cookies
	if error = cookies.SaveCookie(rw, authenticationData.ID, authenticationData.Token); error != nil {
		responses.JSON(rw, http.StatusUnprocessableEntity, responses.APIError{Error: error.Error()})
		return
	}

	// we don't return authentication data. They'll be used only through the cookie
	log.Printf("User %s was successfully authenticated with token %s\n", authenticationData.ID, authenticationData.Token)
	responses.JSON(rw, http.StatusOK, nil)
}
