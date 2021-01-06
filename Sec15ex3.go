//Simple program that illustrates the use of the 'defer' keyword

package main

import (
	"fmt"
)

func main() {
	defer foo()
	fmt.Println("Hello")

}
func foo() {
	fmt.Println("This is inside foo")
}

//OUTPUT:
// Hello
// This is inside foo
