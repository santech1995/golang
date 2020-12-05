//Program to append elements to a slice

package main

import "fmt"

func main() {
	x := []int{4, 6, 5, 67, 887}
	fmt.Println(x)
	x = append(x, 344)
	fmt.Println(x)

	y := []int{324, 546, 9887, 3232}
	x = append(x, y...)
	fmt.Println(x)
}

//OUTPUT:
// [4 6 5 67 887]
// [4 6 5 67 887 344]
// [4 6 5 67 887 344 324 546 9887 3232]
