package main

import "fmt"

// 1 1 2 3 5 8 13 21 34
func fibonacci(n int) int {
	if n <= 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	number := 10
	fmt.Println(fmt.Sprintf("Fibonacci of %d is: %d", number, fibonacci(number)))
}
