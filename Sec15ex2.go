//Program to implement functions in GO

package main

import (
	"fmt"
)

//func parameter identifier parenthesis return_value

func main() {
	f := []int{1, 2, 3, 4, 5, 6, 7}
	n := foo(f...)
	fmt.Println(n)
	b := []int{1, 2, 3, 4, 5, 6, 7}
	n1 := bar(b)
	fmt.Println(n1)
}

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}

	return sum
}

func bar(xi []int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}

	return sum
}

//OUTPUT:
// 28
// 28
