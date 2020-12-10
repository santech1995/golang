//Program to append elements to a map in GO

package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"James": 34,
		"Mary":  21,
	}
	fmt.Println(m)
	//adding elements to the map
	m["Judy"] = 33
	fmt.Println(m)
	for k, v := range m {
		fmt.Println(k, v)
	}
}

//OUTPUT:
// map[James:34 Mary:21]
// map[James:34 Judy:33 Mary:21]
// James 34
// Mary 21
// Judy 33
