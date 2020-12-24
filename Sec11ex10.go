//Program to append an entry into a map in GO

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
	m["fleming_ian"] = []string{"steaks", "cigars", "espionage"}

	delete(m, "no_dr")
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
// This is the record for fleming_ian
//          0 steaks
//          1 cigars
//          2 espionage
