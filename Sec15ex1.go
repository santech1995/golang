//Program to illustrate the use of functions

package main

import (
	"fmt"
)

var x int
var y int

func main() {
	foo()
	foo()
	foo()
	fmt.Println("------------")
	bar()
	bar()
	bar()
}

func foo() int {
	fmt.Println("Inside foo")
	x++
	fmt.Println(x)
	return x
}

func bar() (int, string) {
	s := "Inside bar"
	y++
	fmt.Println(s)
	fmt.Println(y)
	return y, s
}

//OUTPUT:
// Inside foo
// 1
// Inside foo
// 2
// Inside foo
// 3
// ------------
// Inside bar
// 1
// Inside bar
// 2
// Inside bar
// 3
