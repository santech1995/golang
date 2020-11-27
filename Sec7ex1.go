//Simple program to represent a number in decimal,binary and hexadecimal system.

package main

import "fmt"

func main() {
	x := 42
	fmt.Printf("x in decimal = %d\n", x)
	fmt.Printf("x in binary = %b\n", x)
	fmt.Printf("x in hexadecimal = %#x\n", x)
}

//OUTPUT:
// x in decimal = 42
// x in binary = 101010
// x in hexadecimal = 0x2a
