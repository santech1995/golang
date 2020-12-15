//Program to implement slicing in GO

package main

import "fmt"

func main() {
	x := []int{4, 5, 3, 2, 6}
	fmt.Println(x)
	fmt.Println(x[:3])
	fmt.Println(x[1:])
	fmt.Println(x[1:3])

	for i, v := range x {
		fmt.Println(i, v)
	}

	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}

}

//OUTPUT:
// [4 5 3 2 6]
// [4 5 3]
// [5 3 2 6]
// [5 3]
// 0 4
// 1 5
// 2 3
// 3 2
// 4 6
// 0 4
// 1 5
// 2 3
// 3 2
// 4 6
