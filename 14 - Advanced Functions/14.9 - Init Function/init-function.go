package main

import "fmt"

// run before main
// you can have 1 per file (not like main function, 1 per package)
// can be used to initiate a variable

var n int

func init() {
	fmt.Println("Init function")
	n = 10
}

func main() {
	fmt.Println("Main function")
	fmt.Println(n)
}
