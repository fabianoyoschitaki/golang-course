package main

import "fmt"

// there's no method defined in this interface, so everything matches!
func generic(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generic("Test")
	generic(true)
	generic(1)
	generic(float64(1.2345))

	// fmt.Println does that, it receives anything. it needs flexibility to accept any kind of data type
	fmt.Println("Test", true, 1, float64(1.2345))

	// we can, but should not do that, otherwise we create a mess!
	map1 := map[interface{}]interface{}{
		"string": 1,
		true:     false,
	}
	fmt.Println(map1)
}
