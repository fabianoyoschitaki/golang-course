// STARTING PROJECT
// go mod init command-line
// go get github.com/urfave/cli

// RUNNING PROJECT
// to run: go run main.go
// to run ip command: go run main.go ip --host www.amazon.com.br
// to run servers command: go run main.go servers --host www.amazon.com.br

// BUILD PROJECT
// to build: go build
// ./command-line ip --host <host>
// ./command-line servers --host <host>

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
