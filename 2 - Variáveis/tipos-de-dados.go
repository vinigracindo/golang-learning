package main

import (
	"errors"
	"fmt"
)

func main() {
	// int usa a arquitetura do computador.
	// pode ser int8 int16 int32 int64
	var numero int = 1000
	fmt.Println(numero)

	var num2 uint = 100 // Inteiros Naturais (N)

	//alias

	// INT32 = Rune
	var num3 rune = 123456

	// BYTE = UINT8
	var num4 byte = 123
	fmt.Println(num2, num3, num4)
	// end-alias

	var numReal1 float32 = 13.45
	var numReal2 float64 = 15.456

	fmt.Println(numReal1, numReal2)

	var texto int // Toda variável tem o valor inicial. Número é 0 e string é vazio
	fmt.Println(texto)

	var booleano1 bool = true
	fmt.Println(booleano1)

	var erro error // Tipo error. MUITO UTILIZADO EM GOLANG
	fmt.Println(erro)

	var erro2 error = errors.New("Tipo de Erro")
	fmt.Println(erro2)

}
