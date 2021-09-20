package main

import (
	"fmt"
	"strconv"
)

// func that sums all numbers passed
func sum(numbers ...int) int {
	fmt.Println(numbers) // this is a slice!
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

// a fixed parameter + a variadic parameter. It must be the last parameter and there should be only 1 within a function
func fixedParameterWithVariadic(text string, numbers ...int) string {
	if len(numbers) > 0 {
		return text + " First number is " + strconv.Itoa(numbers[0]) + " and Last number is " + strconv.Itoa(numbers[len(numbers)-1])
	}
	return text + " there are no numbers."
}

func main() {
	sum_total := sum(1, 2, 3, 4) // 10
	fmt.Println(sum_total)
	sum_total = sum(1, 2, 3, 4, 6, 7, 8, 9, 10) // 50
	fmt.Println(sum_total)
	sum_total = sum()
	fmt.Println(sum_total) // 0

	// fixed variable + variadic
	fmt.Println(fixedParameterWithVariadic("Fabiano", 1, 2, 3))
	fmt.Println(fixedParameterWithVariadic("Fabiano", 1))
	fmt.Println(fixedParameterWithVariadic("Fabiano"))
}
