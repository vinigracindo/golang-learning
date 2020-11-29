package main

import "fmt"

func somar(x int, y int) int {
	return x + y
}

// Funções podem ter mais de um retorno
func calculosMatematicos(x, y int) (int, int, int, float32) {
	soma := x + y
	subtracao := x - y
	multiplicacao := x * y
	divisao := float32(x) / float32(y)
	return soma, subtracao, multiplicacao, divisao
}

func main() {
	fmt.Println(somar(15, 25))

	var f = func(texto string) string {
		return texto
	}

	fmt.Println(f("Texto"))

	resultadoSoma, resultadoSub, resultadoMul, resultadoDiv := calculosMatematicos(10, 20)
	fmt.Println(resultadoSoma, resultadoSub, resultadoMul, resultadoDiv)

	_, _, multiplicacao, _ := calculosMatematicos(10, 20)
	fmt.Println(multiplicacao)
}
