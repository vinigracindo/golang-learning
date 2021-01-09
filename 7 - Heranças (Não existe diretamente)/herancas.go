package main

import (
	"fmt"
)

// "Herança" em go

type people struct {
	name     string
	lastName string
	age      uint8
}

type student struct {
	people
	course string
}

func main() {
	fmt.Println("Herança")

	p1 := people{"João", "Pedro", 18}
	fmt.Println(p1)

	e1 := student{p1, "Engineer"}
	fmt.Println(e1)
}
