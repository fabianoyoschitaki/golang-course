package main

import (
	"fmt"
)

func main() {
	i := 0
	for i < 10 {
		i++
		fmt.Println("Incrementando i: ", i)
		// time.Sleep(time.Second)
	}
	fmt.Println(i)

	for j := 0; j < 10; j += 2 {
		fmt.Println("Incrementandop j: ", j)
		// time.Sleep(time.Second)
	}
	// fmt.Println(j) // undefined: j

	names := [3]string{"John", "Peter", "Rapha"}
	fmt.Println(names)
	for index, name := range names {
		fmt.Println(index, name)
	}
	// always index first
	for name := range names {
		fmt.Println(name)
	}

	// you dont want the index
	for _, name := range names {
		fmt.Println(name)
	}

	word := "abcdefABCDEF"
	for index, letter := range word {
		fmt.Println(index, letter, string(letter)) // will return int values
	}

	// map for loop
	users := map[int]string{
		1: "One",
		2: "Dois",
	}

}
