package main

import "fmt"

func main() {

	var variavel1 string = "Variável do Tipo String" // Declaração explícita
	variavel2 := "Variável 2"                        // Declaração implícita de variável
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "Variável 3"
		variavel4 string = "Variável 4"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variável 5", "Variável 6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "Constante 1"
	fmt.Println(constante1)

}
