package main

import "fmt"

func main() {

	// creating the slice always points to an array.
	// make creates a 11 position array. it will return a slice that will take first 10 positions of this array
	// internal arrays: type, size, maxSize
	slice := make([]float32, 10, 11)
	fmt.Println(slice)      // [0 0 0 0 0 0 0 0 0 0]
	fmt.Println(len(slice)) // 10
	fmt.Println(cap(slice)) // 11

	// slice doesnt have a max size, so what about the 11?
	slice = append(slice, 5) // len 11, cap 11
	fmt.Println(len(slice))  // 11
	fmt.Println(cap(slice))  // 11

	slice = append(slice, 6) // this will overflow the slice, and will create a new array with doubled size!
	fmt.Println(len(slice))  // 12
	fmt.Println(cap(slice))  // 24

	slice2 := make([]float32, 5)
	fmt.Println(slice2)
	fmt.Println(len(slice2)) // 5
	fmt.Println(cap(slice2)) // 5
	slice2 = append(slice2, 10)
	fmt.Println(len(slice2)) // 6
	fmt.Println(cap(slice2)) // 12
}
