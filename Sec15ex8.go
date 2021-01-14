//Simple program to show how function can return a function

package main

import (
	"fmt"
)

func main() {
	f := foo()
	fmt.Println(f())
}

func foo() func() int {
	return func() int {
		sum := 0
		for i := 0; i < 3; i++ {
			sum += i
		}
		return sum
	}
}

//OUTPUT:
//3
