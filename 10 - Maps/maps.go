package main

import "fmt"

func main() {
	fmt.Println("Maps")

	user := map[string]string{
		"name":     "Vinnicyus",
		"lastName": "Gracindo",
	}

	fmt.Println(user)
	fmt.Println(user["name"])

	// maps aninhados

	user2 := map[string]map[string]string{
		"name": {
			"first": "Vinnicyus",
			"last":  "Gracindo",
		},
		"course": {
			"name":   "Engineer",
			"campus": "UFAL",
		},
	}

	fmt.Println(user2)
	fmt.Println(user2["name"]["first"])

	// remover chaves do map
	delete(user2, "course")
	fmt.Println(user2)

	// adicionar chave
	user2["exams"] = map[string]string{
		"math": "A",
	}

	fmt.Println(user2)

	user["age"] = "19"
	fmt.Println(user)

}
