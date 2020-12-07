//Program to delte elements from a slice in GO

package main

import "fmt"

func main() {
	x := []int{23, 43, 56, 87, 99}
	fmt.Println(x)
	y := []int{121, 232, 345, 456, 768, 978}
	fmt.Println(y)
	x = append(x, y...) //appending the entire y slice to the x slice
	fmt.Println(x)
	x = append(x[:2], x[4:]...) //deleting the x[2] and x[3] elements
	fmt.Println(x)
}

//OUTPUT:
// [23 43 56 87 99]
// [121 232 345 456 768 978]
// [23 43 56 87 99 121 232 345 456 768 978]
// [23 43 99 121 232 345 456 768 978]
