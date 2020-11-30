//Program to convert data of one type to another

package main

import "fmt"

const (
	a     = 42
	b int = 43
)

func main() {
	fmt.Println(a, b)
	fmt.Printf("%T\t%T\n", a, b)
}

//OUTPUT:
// 42 43
// int     int
