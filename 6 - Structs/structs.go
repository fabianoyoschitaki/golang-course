package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	rua    string
	numero uint8
}

func main() {
	fmt.Println("Criando struct usuario")
	// valor 0 struct = valor 0 para todos os campos
	var u usuario
	fmt.Println(u)
	u.idade = 33
	u.nome = "Fabiano"
	fmt.Println(u)

	// inferencia de tipo para struct
	u2 := usuario{"Joao", 10, endereco{"Av Paulista", 100}}
	fmt.Println(u2)

	// inferencia de tipos, mas sem todos os valores
	u3 := usuario{idade: 1, endereco: endereco{rua: "Av Paulista"}}
	fmt.Println(u3)
}
