package main

import "fmt"

func main() {

	// explicito
	var variavel1 string = "Variavel 1"
	fmt.Println(variavel1)

	// implicito, inferencia de tipo
	variavel2 := "Variavel 2"
	fmt.Println(variavel2)

	// varias ao mesmo tempo
	var (
		variavel3 string = "lalala"
		variavel4 int    = 1
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "lelele", 2
	fmt.Println(variavel5, variavel6)

	// constantes
	const constante1 string = "Eu sou uma constante"
	fmt.Println(constante1)
	//constante1 = "Nao posso mudar"

	// inverter 2 variaveis
	variavel1, variavel2 = variavel2, variavel1
}
