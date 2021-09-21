// go mod init command-line
// go get github.com/urfave/cli
// to run: go run main.go

package main

import (
	"command-line/app"
	"fmt"
	"log"
	"os"
)

// it will only call app.go, where it's our actual application
func main() {
	fmt.Println("Starting point")
	application := app.Generate()
	// error := application.Run(os.Args)
	// if error != nil {
	// 	log.Fatal(error)
	// }

	// IF INIT
	if error := application.Run(os.Args); error != nil {
		log.Fatal(error)
	}
}
