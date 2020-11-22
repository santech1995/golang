//Simple program to illiustrate the use of short declaration operator

package main

import "fmt"

func main() {
	x := 20
	fmt.Println(x)
	x = 30
	fmt.Println(x)
	y := 100 + 20
	fmt.Println(y)
}

//OUTPUT:
// 20
// 30
// 120
