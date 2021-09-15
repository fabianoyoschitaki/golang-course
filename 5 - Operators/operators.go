package main

import "fmt"

func main() {
	//-------------------------------------
	// ARITMETICOS +, -, *, /
	//-------------------------------------
	soma := 1 + 1
	subtracao := 4 - 1
	multiplicacao := 5 * 5
	divisao := 10 / 2
	restoDivisao := 7 % 4
	fmt.Println(soma, subtracao, multiplicacao, divisao, restoDivisao)

	// Nao pode fazer nada com variaveis de tipos diferentes. Go eh fortemente tipado
	var numero16 int16 = 10
	var numero32 int32 = 10
	somaTiposDiferentes := int32(numero16) + numero32
	fmt.Println(somaTiposDiferentes)

	//-------------------------------------
	// ATRIBUICAO
	//-------------------------------------
	var variavel string = "Operador de atribuicao"
	variavel2 := "Inferencia de Tipo"
	fmt.Println(variavel, variavel2)

	//-------------------------------------
	// OPERADORES RELACIONAIS
	//-------------------------------------
	fmt.Println(1 > 2)
	fmt.Println(1 == 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	//-------------------------------------
	// OPERADORES LOGICOS
	//-------------------------------------
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && verdadeiro && true)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)

	//-------------------------------------
	// OPERADORES UNARIOS
	//-------------------------------------
	n := 10
	n++ //pos fix, there's no pre fix
	n--
	n += 2
	n -= 1
	n *= 2
	n /= 2
	n %= 2
	fmt.Println(n)

	//-------------------------------------
	// OPERADORES TERNARIOS
	//-------------------------------------
	// nao existe em Go. A premissa eh ter 1 forma de fazer as coisas. if else
	if 5 > 6 {

	} else {

	}
}
