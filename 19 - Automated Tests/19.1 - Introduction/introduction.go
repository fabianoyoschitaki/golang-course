package main

import (
	"fmt"
	"introduction-tests/addresses"
)

// remember to run: go mod init introduction-tests before

// go run introduction.go
func main() {
	typeAddress := addresses.TypeOfAddress("Avenue Paulista")
	fmt.Println(typeAddress)
}
