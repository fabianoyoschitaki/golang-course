package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	fmt.Println("Running webapp")

	// loading configuration (environment variables)
	fmt.Println("Loading environment variables...")
	config.LoadConfiguration()
	fmt.Println("Environment variables successfully loaded!")

	// loading secureCookie
	fmt.Println("Loading secureCookie...")
	cookies.Configure(config.HashKey, config.BlockKey)
	fmt.Println("secureCookie successfully loaded!")

	// loading HTML templates
	fmt.Println("Loading templates...")
	utils.LoadTemplates()
	fmt.Println("Templates successfully loaded!")

	// loading HTTP routes
	fmt.Println("Generating routes...")
	r := router.Generate()
	fmt.Println("Routes successfully generated!")

	fmt.Printf("Listening at port %d\n", config.ApplicationPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.ApplicationPort), r))
}
