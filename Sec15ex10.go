//Program illustrating the closure property in GO

package main

import (
	"fmt"
)

func main() {
	f := foo()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

//closure is enclosing a variable within a block to restrict its scope

func foo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

//OUTPUT:
// 1
// 2
// 3
// 4
