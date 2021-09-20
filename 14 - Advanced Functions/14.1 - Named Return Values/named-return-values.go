package main

import "fmt"

// named returns
func calc(n1, n2 int) (sum int, subtraction int) {
	// does not need := since it's already declared
	sum = n1 + n2
	subtraction = n1 - n2

	// does not need to return the variables since it knows what to do
	return
}

func main() {
	sum, subtraction := calc(1, 2)
	fmt.Println(sum, subtraction)
}
