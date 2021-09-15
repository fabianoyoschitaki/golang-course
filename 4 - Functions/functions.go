package main

import "fmt"

func somar(n1 int8, n2 int8) int16 {
	return int16(n1) + int16(n2)
}

func main() {
	soma := somar(127, 127)
	fmt.Println(soma)

	var printAnything = func(anything string) string {
		fmt.Println("I print " + anything)
		return "I print " + anything
	}

	var resultado string = printAnything("Fabiano")
	fmt.Println(resultado)

	r1, r2, r3, r4 := calculosMatematicos(2, 2)
	fmt.Println("Soma eh ", r1)
	fmt.Println("Subtracao eh ", r2)
	fmt.Println("Multiplicacao eh ", r3)
	fmt.Println("Divisao eh ", r4)

	// caso nao queira usar um retorno, usar _
	s, _, _, _ := calculosMatematicos(10, 10)
	fmt.Println(s)
}

// funcoes podem ter mais de 1 retorno! (se parametros forem do mesmo tipo, n precisa declarar 1 a 1)
func calculosMatematicos(n1, n2 int8) (int8, int8, int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	multiplicacao := n1 * n2
	divisao := n1 / n2
	return soma, subtracao, multiplicacao, divisao
}
