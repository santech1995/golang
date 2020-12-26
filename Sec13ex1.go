//Program to declare a struct and create objects of the struct

package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	flavours  []string
}

func main() {

	p1 := person{
		firstName: "James",
		lastName:  "Bond",
		flavours: []string{
			"strawberry",
			"vanilla",
		},
	}
	p2 := person{
		firstName: "Miss",
		lastName:  "Briganza",
		flavours: []string{
			"chocolate",
			"blackcurrent",
		},
	}
	fmt.Println(p1.firstName)
	fmt.Println(p1.lastName)
	for i, v := range p1.flavours {
		fmt.Println(i, v)
	}
	fmt.Println(p2.firstName)
	fmt.Println(p2.lastName)
	for i, v := range p2.flavours {
		fmt.Println(i, v)
	}
}

//OUTPUT:
// James
// Bond
// 0 strawberry
// 1 vanilla
// Miss
// Briganza
// 0 chocolate
// 1 blackcurrent
