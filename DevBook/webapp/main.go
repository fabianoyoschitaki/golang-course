package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	fmt.Println("Running webapp")

	fmt.Println("Loading templates...")
	utils.LoadTemplates()
	fmt.Println("Templates successfully loaded!")

	fmt.Println("Generating routes...")
	r := router.Generate()
	fmt.Println("Routes successfully generated!")

	port := ":3000"
	fmt.Printf("Listening at port %s", port)
	log.Fatal(http.ListenAndServe(port, r))
}
