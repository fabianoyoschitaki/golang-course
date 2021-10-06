package controllers

import "net/http"

// CreateUser creates a new user
func CreateUser(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Creating user"))
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
