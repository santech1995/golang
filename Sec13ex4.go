//Program that implements a struct,map and slice in GO

package main

import (
	"fmt"
)

func main() {
	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Mini": 333,
			"Q":    222,
			"M":    888,
		},
		favDrinks: []string{
			"Coke",
			"Water",
		},
	}
	fmt.Println(s.first)
	fmt.Println(s.friends)
	fmt.Println(s.favDrinks)

	for k, v := range s.friends {
		fmt.Println(k, v)
	}

	for i, val := range s.favDrinks {
		fmt.Println(i, val)
	}
}

//OUTPUT:
// James
// map[M:888 Mini:333 Q:222]
// [Coke Water]
// Mini 333
// Q 222
// M 888
// 0 Coke
// 1 Water
