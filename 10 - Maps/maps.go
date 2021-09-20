package main

import "fmt"

func main() {
	// key, value
	people := map[string]string{
		"name":     "John",
		"lastName": "Doe",
	}
	fmt.Println(people)         // map[lastName:Doe name:John]
	fmt.Println(people["name"]) // John

	// map of map
	mapOfMap := map[string]map[string]string{
		"user1": {
			"name":     "John",
			"lastName": "Doe",
		},
		"user2": {
			"age":    "10",
			"height": "100",
		},
	}
	fmt.Println(mapOfMap) // map[user1:map[lastName:Doe name:John] user2:map[age:10 height:100]]
	// remove from map
	delete(mapOfMap, "user1")
	fmt.Println(mapOfMap) // map[user2:map[age:10 height:100]]

	mapOfMap["user1"] = map[string]string{
		"name":     "John",
		"lastName": "Doe",
	}
	fmt.Println(mapOfMap)
}
