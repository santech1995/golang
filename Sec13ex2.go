//Program which combines the implementation of map,struct and slice
//in GO

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
	m := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}
	fmt.Println(m)

	for _, v := range m {

		fmt.Println(v.firstName)
		fmt.Println(v.lastName)
		for i, val := range v.flavours {
			fmt.Println(i, val)
		}
		fmt.Println("--------------------------")
	}
}

//OUTPUT:
// map[Bond:{James Bond [strawberry vanilla]} Briganza:{Miss Briganza [chocolate blackcurrent]}]
// James
// Bond
// 0 strawberry
// 1 vanilla
// --------------------------
// Miss
// Briganza
// 0 chocolate
// 1 blackcurrent
// --------------------------
