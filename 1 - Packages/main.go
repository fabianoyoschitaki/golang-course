package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("valid@gmail.com")
	fmt.Println(erro)

	erro2 := checkmail.ValidateFormat("invalid@@gmail.com")
	fmt.Println(erro2)
}
