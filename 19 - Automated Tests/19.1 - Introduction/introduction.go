package main

import (
	"fmt"
	"introduction-tests/addresses"
)

// #IMPORTANT
// remember to run: go mod init introduction-tests before

// #IMPORTANT
// you can only have 1 package per folder, the exception is for tests.

// go run introduction.go

// instead of going to each package and run: go test, you can go to root folder (like where this file is in) and run: go test ./...
//	?		introduction-tests      [no test files]
//	ok      introduction-tests/addresses    0.006s

// #IMPORTANT
// you can run verbose mode: go test -v

// #IMPORTANT
// you can check test coverage: go test --cover

// #IMPORTANT
// to check what's not covered: go test --coverprofile coverage.txt
// to understand coverage.txt: go tool cover --func=coverage.txt
// to understand which lines are not covered: go tool cover --html=coverage.txt
func main() {
	typeAddress := addresses.TypeOfAddress("Avenue Paulista")
	fmt.Println(typeAddress)
}
