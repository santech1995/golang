//Program to implement the comma-ok idiom in GO

package main

import (
	"fmt"
)

func main() {
	//deleting an element from a map:
	//delete(<map_name>,"key")
	m := map[string]int{
		"James": 34,
		"Mary":  21,
	}
	fmt.Println(m)
	delete(m, "James")
	fmt.Println(m)

	//comma-ok idiom
	if v, ok := m["Mary"]; ok {
		fmt.Println("value:", v)
		fmt.Println("exists or not:", ok)
	}
}

//OUTPUT:
// map[James:34 Mary:21]
// map[Mary:21]
// value: 21
// exists or not: true
