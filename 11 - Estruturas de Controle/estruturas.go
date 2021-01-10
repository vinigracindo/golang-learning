package main

import "fmt"

func main() {
	fmt.Println("Estruturas de Controle. SEMELHANTE AS OUTRAS LINGUAGENS")

	// diferença
	// if init (inicializa varíavel dentro do escopo do IF)

	if numero := 10; numero > 0 {
		fmt.Println("O número é positivo")
	} else {
		fmt.Println("O número é negativo")
	}
}
