//Simple program to illustrate the use of pointers

package main

import (
	"fmt"
)

func main() {
	x := 42
	fmt.Println("x =", x)
	fmt.Println("address of x:", &x)
}

//OUTPUT:
// x = 42
// address of x: 0xc0000a2058
