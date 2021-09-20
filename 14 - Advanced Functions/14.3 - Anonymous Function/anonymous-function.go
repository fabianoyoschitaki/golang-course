package main

import "fmt"

func main() {

	text := "Fabiano"
	result := func(text string) string {
		return fmt.Sprintf("I'm an anonymous function: %s has length %d", text, len(text))
	}(text) // calling the anonymoys function with parameter

	// return of the anonymous function
	fmt.Println(result)
}
