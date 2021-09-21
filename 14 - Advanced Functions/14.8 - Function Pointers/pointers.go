package main

import "fmt"

func reverseSignByCopy(n int) int {
	return n * -1
}

func reverseSignByReference(n *int) {
	// dereferencing
	*n = *n * -1
}

func main() {
	numberByCopy := 10
	resultByCope := reverseSignByCopy(numberByCopy)
	fmt.Println(numberByCopy) // 10
	fmt.Println(resultByCope) // -10

	// let's pass by reference
	numberByReference := 5
	reverseSignByReference(&numberByReference)
	fmt.Println(numberByReference) // it will change to -5
}
