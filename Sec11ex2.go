//Program in GO to iterate through a slice and determine its type using type specifier

package main

import (
	"fmt"
)

func main() {
	x := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("type of slice is %T", x)
}

//OUTPUT:-
// 0 10
// 1 20
// 2 30
// 3 40
// 4 50
// 5 60
// 6 70
// 7 80
// 8 90
// 9 100
// type of slice is []int
