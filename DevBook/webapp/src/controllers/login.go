package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
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

	response, error := http.Post("http://localhost:5000/login", "application/json", bytes.NewBuffer(login))
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
