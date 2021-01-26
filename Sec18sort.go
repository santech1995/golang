//Simple program to implement the sort method from Sort package
//to sort a slice of integers and a slice of strings

package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{6, 5, 8, 3, 4, 1, 2}
	xs := []string{"Zebra", "Giraffe", "Kitten", "Ant", "Buffalo"}

	fmt.Println("Integers before sorting:", xi)
	sort.Ints(xi)
	fmt.Println("Strings after sorting:", xi)

	fmt.Println("Strings before sorting:", xs)
	sort.Strings(xs)
	fmt.Println("Strings after sorting:", xs)
}

//OUTPUT:
// Integers before sorting: [6 5 8 3 4 1 2]
// Strings after sorting: [1 2 3 4 5 6 8]
// Strings before sorting: [Zebra Giraffe Kitten Ant Buffalo]
// Strings after sorting: [Ant Buffalo Giraffe Kitten Zebra]
