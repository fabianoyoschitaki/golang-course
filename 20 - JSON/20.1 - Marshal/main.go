package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type Dog struct {
	Name  string `json:"name"` // keys when turning into JSON
	Breed string `json:"breed"`
	Age   uint   `json:"age"`
}

func main() {
	//------------------------------------------------------------------
	// STRUCT to JSON
	//------------------------------------------------------------------
	dog := Dog{"Rex", "Normal Dog", 5}
	dogJSON, error := json.Marshal(dog)
	if error != nil {
		// stops execution
		log.Fatal(error)
	}

	// returns slice of uint8:
	// [123 34 110 97 109 101 34 58 34 82 101 120 34 44 34 98 114 101 101 100 34 58 34 78 111 114 109 97 108 32 68 111 103 34 44 34 97 103 101 34 58 53 125]
	fmt.Println(dogJSON)

	// how to make it as JSON? using another go package:
	// {"name":"Rex","breed":"Normal Dog","age":5}
	fmt.Println(bytes.NewBuffer(dogJSON))

	//------------------------------------------------------------------
	// MAP to JSON
	//------------------------------------------------------------------
	dog2 := map[string]string{
		"name":  "Tobby",
		"breed": "Poodle",
	}
	dog2JSON, error := json.Marshal(dog2)
	if error != nil {
		log.Fatal(error)
	}
	// [123 34 98 114 101 101 100 34 58 34 80 111 111 100 108 101 34 44 34 110 97 109 101 34 58 34 84 111 98 98 121 34 125]
	fmt.Println(dog2JSON)
	// {"breed":"Poodle","name":"Tobby"}
	fmt.Println(bytes.NewBuffer(dog2JSON))
}
