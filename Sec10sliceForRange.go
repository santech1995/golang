//Program to range over the elements of a slice in GO

package main

import "fmt"

func main() {
	x := []int{4, 5, 6, 7, 8}
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])
	fmt.Println(len(x))
	for i, v := range x {
		fmt.Println(i, v)
	}
}

//OUTPUT:
// 4
// 5
// 6
// 7
// 8
// 5
// 0 4
// 1 5
// 2 6
// 3 7
// 4 8
