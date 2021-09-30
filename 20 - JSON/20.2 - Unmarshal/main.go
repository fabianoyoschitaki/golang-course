package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Dog struct {
	Name  string `json:"name"` // keys when turning into JSON
	Breed string `json:"-"`    //`json:"breed"` will make breed not to be marshaled into the struct: {Rex  5}
	Age   uint   `json:"age"`
}

func main() {
	//------------------------------------------------------------------
	// JSON to STRUCT
	//------------------------------------------------------------------
	dogJSON := `{"name":"Rex","breed":"Normal Dog","age":5}`

	// type inference
	// d := Dog{}
	var dog Dog

	// unmarshal receives slice of bytes and the memory address of the object we want to unmsarshal into
	if error := json.Unmarshal([]byte(dogJSON), &dog); error != nil { // if init
		log.Fatal(error)
	}
	fmt.Println(dog)

	//------------------------------------------------------------------
	// JSON to MAP
	//------------------------------------------------------------------
	dog2JSON := `{"breed":"Poodle","name":"Tobby"}`
	var dog2 map[string]string
	if error := json.Unmarshal([]byte(dog2JSON), &dog2); error != nil { // if init
		log.Fatal(error)
	}
	fmt.Println(dog2)
}
