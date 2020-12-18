//Program to iterate through a array in GO

package main

import (
	"fmt"
)

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x[:5]) //starting from beginning thru' upto but not inc
	// 5th position
	fmt.Println(x[5:])  //inclusing 5th pos to the end
	fmt.Println(x[2:7]) //starting from and inc 2nd upto & not inc 7th
	fmt.Println(x[1:6]) //starting from and inc 1st upto & not inc 6th

}

//OUTPUT:-
// [42 43 44 45 46]
// [47 48 49 50 51]
// [44 45 46 47 48]
// [43 44 45 46 47]
