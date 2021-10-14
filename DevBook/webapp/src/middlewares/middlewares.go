package middlewares

import (
	"log"
	"net/http"
	"webapp/src/cookies"
)

// Logger logs request data
func Logger(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunction(rw, r)
	}
}

// Authenticate verifies if the cookie DEVBOOK_DATA exists, it doesn't validate if the ID and/or token are valid
func Authenticate(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// values, error := cookies.ReadCookie(r)
		// fmt.Println(values, error)

		// #IMPORTANT
		// we just want to make sure we could read the cookie, not validating its content. the backend API who will validate the Bearer token
		// then we redirect to /login
		if _, error := cookies.ReadCookie(r); error != nil {
			log.Printf("Error on %s %s %s (will be redirected to /login): %s", r.Method, r.RequestURI, r.Host, error)
			http.Redirect(rw, r, "/login", http.StatusMovedPermanently)
			return
		}
		log.Printf("Authentication on %s %s %s successful!", r.Method, r.RequestURI, r.Host)
		nextFunction(rw, r)
	}
}
