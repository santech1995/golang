//Simple program to assign a function to a variable

package main

import (
	"fmt"
)

func main() {
	f := foo()
	fmt.Println("The number returned from the function assigned is: ", f)
	fmt.Printf("%T\n", f)
}

func foo() int {
	x := 32
	return x
}

//OUTPUT:
// The number returned from the function assigned is:  32
//int
