//Program to implement a map and iterate through it in GO

package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		"bond_james":      []string{"shaken,not stirred", "Martinis", "women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}
	//fmt.Println(m)

	for k, v1 := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v1 {
			fmt.Println("\t", i, v2)
		}
	}

}

//OUTPUT:
// This is the record for bond_james
//          0 shaken,not stirred
//          1 Martinis
//          2 women
// This is the record for moneypenny_miss
//          0 James Bond
//          1 Literature
//          2 Computer Science
// This is the record for no_dr
//          0 Being evil
//          1 Ice cream
//          2 Sunsets
