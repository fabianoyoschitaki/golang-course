package main

import "fmt"

func main() {

	// regular if/else
	number := 10
	if number > 15 {
		fmt.Println("Bigger")
	} else {
		fmt.Println("Equals or smaller")
	}

	// if init. you limit it to the scope of the if else
	if number2 := number; number2 > 20 {
		fmt.Println("Bigger")
	} else if number2 > 1 {
		fmt.Println("Equals or smaller")
	}
	// number2 can't be accessed

}
