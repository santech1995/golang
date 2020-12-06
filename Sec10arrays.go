//Program to implement usage of arrays in GO

package main

import "fmt"

func main() {
	var x [5]int
	fmt.Println(x)
	x[2] = 21
	fmt.Println(x)
	fmt.Println(len(x))
}

//OUTPUT:
// [0 0 0 0 0]
// [0 0 21 0 0]
// 5
