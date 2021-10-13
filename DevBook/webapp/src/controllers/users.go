package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"webapp/src/responses"
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
	response, error := http.Post("http://localhost:5000/users", "application/json", bytes.NewBuffer(userToCreate))
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
