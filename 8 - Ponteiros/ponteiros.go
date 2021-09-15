package main

import "fmt"

// Ponteiro salva endereco de memoria

func main() {

	var variavel1 int = 10
	// same values (copy)
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	// this only changes variavel1
	variavel1++
	fmt.Println(variavel1, variavel2)

	// pointer is a memory reference
	var variavel3 int
	var pointer *int

	variavel3 = 100
	pointer = &variavel3

	fmt.Println(variavel3, pointer)

	variavel3++
	fmt.Println(variavel3, pointer, *pointer) // *pointer = Dereferencing
}
