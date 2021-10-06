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
	fmt.Println(config.ApplicationPort)
	fmt.Println(config.DatabaseConnectionString)

	fmt.Println("Running API")
	r := router.Generate()

	// starting the server at port 5000
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.ApplicationPort), r))
}
