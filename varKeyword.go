//Simple program to illustrate the use of the 'var' keyword

package main

import "fmt"

var y int

func main() {
	x := 23
	fmt.Println(x)
	foo()
	y = 20
	fmt.Println("y again:", y)
}
func foo() {
	fmt.Println(y)
}

//OUTPUT:
// 23
// 0
// y again: 20
