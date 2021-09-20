package main

import (
	"fmt"
)

// 0 1 1 2 3 5 8 13 21 34
func fibonacci(n uint) uint {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	number := uint(10)
	fmt.Println(fmt.Sprintf("Fibonacci of %d is: %d", number, fibonacci(number)))

	// printing all fibo
	for i := uint(0); i <= number; i++ {
		fmt.Println(fmt.Sprintf("Fibo of position %d is %d", i, fibonacci(i)))
	}
}
