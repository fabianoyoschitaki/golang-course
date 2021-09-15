package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

// estudante "herda" de pessoa
type estudante struct {
	// "heranca em go", nao especifica o tipo
	pessoa
	p         pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Heranca")

	p1 := pessoa{"Joao", "Pedro", 20, 170}
	fmt.Println(p1)

	estudante := estudante{p1, pessoa{"Outra", "Pessoa", 1, 1}, "Engenharia", "USP"}
	fmt.Println(estudante)
	// nao precisamos acessar estudante.pessoa.nome, a nao ser que tivessemos declarado o tipo
	fmt.Println(estudante.nome)
	fmt.Println(estudante.p.nome)
}
