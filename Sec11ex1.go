//Program to display an array of integers

package main

import (
	"fmt"
)

func main() {
	x := [5]int{10, 20, 30, 40, 50}

	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("Type of array is %T", x)
}

//OUTPUT:
//0 10
//1 20
//2 30
//3 40
//4 50
//Type of array is [5]int  here size/length is a part of the type of array
