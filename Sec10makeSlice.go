//Program to create a slice and append elements to it

package main

import "fmt"

func main() {
	x := []int{21, 32, 34, 54, 65, 78}
	fmt.Println(x)
	y := []int{121, 223, 334, 445, 665, 667}
	fmt.Println(y)
	//format of make():
	//make([]slice_type,length,capacity)
	z := make([]int, 10, 10)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))
	z[0] = 21
	z[9] = 31
	z = append(z, 32)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z)) //when size is greater than old capacity,the
	//value of capacity changes @ runtime new capacity = 2 * old capacity
	//old capacity = 10 new capacity = 20
}

//OUTPUT:
// [21 32 34 54 65 78]
// [121 223 334 445 665 667]
// [0 0 0 0 0 0 0 0 0 0]
// 10
// 10
// [21 0 0 0 0 0 0 0 0 31 32]
// 11
// 20
