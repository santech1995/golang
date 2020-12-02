//Program to represent a number in decimal,binary and hexadecimal

package main

import "fmt"

func main() {
	a := 54
	fmt.Printf("a = %d\t%b\t%#x\n", a, a, a)
	b := a << 1
	fmt.Printf("b = %d\t%b\t%#x", b, b, b)

}

//OUTPUT:
// a = 54  110110  0x36
// b = 108 1101100 0x6c
