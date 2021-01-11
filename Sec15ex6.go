//Simple program to illustrate the use of anonymous function

package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("inside anonymous function")
	}()
	fmt.Println("outside anonymous function")
}

//OUTPUT:
// inside anonymous function
// outside anonymous function
