//Program to implement variadic parameter in GO

package main

import (
	"fmt"
)

func main() {
	foo(2, 3, 4, 5, 6, 7, 8, 9)
}

//(...int) is the variadic parameter because it accepts any number
//of integers
func foo(x ...int) {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}

//OUTPUT:
// [2 3 4 5 6 7 8 9]
// []int
