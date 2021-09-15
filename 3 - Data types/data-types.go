package main

import (
	"errors"
	"fmt"
)

func main() {
	// 4 tipos de inteiros: int8, int16, int32, int64

	// -128 a 127
	var i8 int8 = 127
	i8 = -128
	fmt.Println(i8)

	// -32768 a 32767
	var i16 int16 = 32767
	i16 = -32768
	fmt.Println(i16)

	// -2147483648 a 2147483647
	var i32 int32 = 2147483647
	i32 = -2147483648
	fmt.Println(i32)

	// -9223372036854775808 a 9223372036854775807
	var i64 int64 = 9223372036854775807
	i64 = -9223372036854775808
	fmt.Println(i64)

	// int sozinho vai usar a arquitetura do computador como base
	var i int = 9223372036854775807
	fmt.Println(i)

	// unsigned int, apenas positivos
	var ui8 uint8 = 255
	fmt.Println(ui8)

	// alias para int32: rune
	var i32rune rune = 2147483647
	i32rune = -2147483648
	fmt.Println(i32rune)

	// alias para uint8: byte
	var i8byte byte = 255
	i8byte = 0
	fmt.Println(i8byte)

	// numeros reais: float32 e float64
	var f32 float32 = 1.490812049
	fmt.Println(f32)

	// numeros reais: float32 e float64
	var f64 float64 = 1.32890389021389082938109
	fmt.Println(f64)

	// 32 ou 64 depende da arquitetura
	float3 := 3.121931371737137103
	fmt.Println(float3)

	// string
	var s1 string = "Eu sou uma string"
	fmt.Println(s1)

	s2 := "Eu tbm sou uma string"
	fmt.Println(s2)

	// no go nao tem char! vai printar 66, da tabela ASCII
	// vai ser do tipo int. Por isso o rune
	char := 'B'
	fmt.Println(char)

	// valor 0: valor atribuido a uma variavel se vc nao inicializar
	// string = ""
	// int = 0
	// err = <nil>
	var s3 string
	fmt.Println(s3)

	// booleano, valor 0 = false
	var bool1 bool = true
	bool1 = false
	fmt.Println(bool1)

	// error: eh um tipo no go! muito usado em Go.
	var erro1 error = errors.New("Erro interno")
	fmt.Println(erro1)
}
