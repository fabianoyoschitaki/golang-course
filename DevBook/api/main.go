package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadConfiguration()
	// fmt.Println(config.ApplicationPort)
	// fmt.Println(config.DatabaseConnectionString)

	r := router.Generate()

	// starting the server at port 5000
	fmt.Printf("Running API at port %d", config.ApplicationPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.ApplicationPort), r))
}
