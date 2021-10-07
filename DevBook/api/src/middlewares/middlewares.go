package middlewares

import (
	"api/src/authentication"
	"api/src/responses"
	"log"
	"net/http"
)

// middlewares contains layers between request and response. Often used to something that should be appied to ALL routes

// Authenticate verifies if user making the request is authenticated
func Authenticate(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if error := authentication.ValidateToken(r); error != nil {
			responses.Error(rw, http.StatusUnauthorized, error)
			return
		}
		nextFunction(rw, r)
	}
}

// Logger logs request data
func Logger(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunction(rw, r)
	}
}
