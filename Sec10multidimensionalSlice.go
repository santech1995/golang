//Program to implement multidimensional arrays

package main

import "fmt"

func main() {
	veggy := []string{"tomato", "carrot", "beans", "potato"}
	fmt.Println(veggy)
	fruity := []string{"banana", "kiwi", "apple", "mango"}
	fmt.Println(fruity)
	vegFruit := [][]string{veggy, fruity} //a slice of slice = 2D slice
	fmt.Println(vegFruit)
}

//OUTPUT:
// [tomato carrot beans potato]
// [banana kiwi apple mango]
// [[tomato carrot beans potato] [banana kiwi apple mango]]
