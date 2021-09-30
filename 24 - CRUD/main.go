package main

import (
	"basic-crud/server"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// #IMPORTANT go get github.com/gorilla/mux
func main() {

	// configure routes
	router := mux.NewRouter()
	router.HandleFunc("/users", server.CreateUser).Methods(http.MethodPost) // POST route

	fmt.Println("Listening at :5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
