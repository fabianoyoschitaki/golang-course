package main

import (
	"fmt"
	"reflect"
)

func main() {

	// ARRAYS
	// all of the same type. size is mandatory and it's fixed
	var array [5]int
	array[0] = 10
	fmt.Println(array)

	array2 := [5]string{"1", "2", "3", "4", "5"}
	fmt.Println(array2)
	// can't add more than what you created!
	// array2[5] = "6"

	// ... will create the size you provide, it doesnt make it dynamic
	array3 := [...]int{1, 2, 3}
	fmt.Println(array3)

	// ---------------------
	// SLICES
	var slice []int = []int{1, 2, 3, 4}
	fmt.Println(slice)

	// proof arrays and slices are different
	// []int and [5]int
	fmt.Println(reflect.TypeOf(slice), reflect.TypeOf(array))

	// add new values
	slice = append(slice, 1000)
	fmt.Println(slice)

	// we can create a slice from an array. [inclusive, exclusive] it's a pointer
	array4 := [...]string{"a", "b", "c"}
	fmt.Println(array4)
	slice2 := array4[0:2]
	fmt.Println(slice2)

	// changing the array will modify the slice
	array4[0] = "X"
	fmt.Println(slice2)

}
