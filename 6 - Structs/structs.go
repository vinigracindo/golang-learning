package main

import (
	"fmt"
)

type userModel struct {
	name string
	age  uint8
}

func main() {
	var user userModel
	user.name = "Vinnicyus Gracindo"
	user.age = 31
	fmt.Println(user)

	user2 := userModel{"José Abreu", 58}
	fmt.Println(user2)

	user3 := userModel{age: 48}
	fmt.Println(user3)
}
