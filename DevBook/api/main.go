package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

// #IMPORTANT this is only a one time function to generate a random base64 secret for JWT
// func init() {
// 	key := make([]byte, 64)
// 	if _, error := rand.Read(key); error != nil {
// 		log.Fatal(error)
// 	}

// 	base64Key := base64.StdEncoding.EncodeToString(key)
// 	fmt.Printf("JWT secret key is: %s", base64Key)
// }

func main() {
	config.LoadConfiguration()
	// fmt.Println(config.ApplicationPort)
	// fmt.Println(config.DatabaseConnectionString)

	r := router.Generate()

	// starting the server at port 5000
	log.Printf("Running API at port %d\n", config.ApplicationPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.ApplicationPort), r))
}
