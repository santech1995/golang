//Program to implement nested for range loops in GO

package main

import (
	"fmt"
)

func main() {
	x := []string{"James", "Bond", "Shaken,not stirred"}
	fmt.Println(x)
	y := []string{"Miss", "Moneypenny", "Hellooooo,James"}
	fmt.Println(y)
	z := [][]string{x, y}
	//for range within for range
	for i, v := range z {
		fmt.Println("record:", i)
		for j, val := range v {
			fmt.Printf("\t index position: %v \t value: \t %v\n", j, val)
		}
	}
}

//OUTPUT:-
// [James Bond Shaken,not stirred]
// [Miss Moneypenny Hellooooo,James]
// record: 0
//          index position: 0       value:          James
//          index position: 1       value:          Bond
//          index position: 2       value:          Shaken,not stirred
// record: 1
//          index position: 0       value:          Miss
//          index position: 1       value:          Moneypenny
//          index position: 2       value:          Hellooooo,James
