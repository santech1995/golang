//Simple program to illustrate the use custom data type 'hotdog'

package main

import "fmt"

type hotdog int

var x hotdog
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("Custom data type is %T\n", x)
	x = 42
	fmt.Println(x)
	y := int(x)
	fmt.Println(y)
	fmt.Printf("Built-in data type is %T\n", y)
}

//OUTPUT:
// 0
// Custom data type is main.hotdog
// 42
// 42
// Built-in data type is int
