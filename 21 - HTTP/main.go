package main

import (
	"log"
	"net/http"
)

func main() {
	// HTTP is a communication protocol - base of web - client server style
	// Routes comprised of 1. URI - universal resource identifier, 2. Method - GET POST PUT DELETE etc

	// create a route (w - writer, r - request). Declare routes before
	http.HandleFunc("/home", home)

	http.HandleFunc("/users", users)

	// create a HTTP server
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Users page..."))
}
