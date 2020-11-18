//Simple program to create own datatype called 'hotdog' which has an
//underlying type of 'int'

package main

import "fmt"

var a int

type hotdog int

var b hotdog

func main() {
	a = 42
	fmt.Println("\na =", a)
	fmt.Printf("a is of type %T\n", a)
	fmt.Println("")
	b = 30
	fmt.Println("")
	fmt.Println("b =", b)
	fmt.Printf("b is of type %T\n", b)
}
