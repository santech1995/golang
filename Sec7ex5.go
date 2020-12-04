//Program to create a variable of type string using a raw string literal

package main

import "fmt"

func main() {
	a := `this is 
	a 
	raw string literal
	"you see?"
	another thing`
	fmt.Println(a)
}

//OUTPUT:
// this is
//         a
//         raw string literal
//         "you see?"
//         another thing
