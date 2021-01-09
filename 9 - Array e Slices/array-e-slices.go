package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Array e Slices")

	/**************************
	*					ARRAY						*
	**************************/
	var array1 [5]int
	array1[0] = 15
	array1[4] = 10
	fmt.Println(array1)

	array2 := [5]string{"Posição1", "Posição2"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(array3)

	/**************************
	*					SLICE						*
	**************************/
	slice := []int{10, 11, 12, 13, 14, 15}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(array3))
	fmt.Println(reflect.TypeOf(slice))

	slice = append(slice, 20)
	fmt.Println(slice)

	slice2 := array3[0:3]
	fmt.Println(slice2)

	/**************************
	*			ARRAYS INTERNOS			*
	**************************/

	slice3 := make([]float32, 10, 11) //Tamanho e Capacidade
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6) // Quando estoura o limite, adiciona 1 e dobra a capacidade

	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]float32, 5) // Quando não especifica capacidade, capacidade=tamanho
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

}
