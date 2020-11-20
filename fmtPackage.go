//Simple program that explores the fmt package

package main

import "fmt"

var y = 42

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)  //tells the type of data in y
	fmt.Printf("%b\n", y)  //gives the binary equivalent of y
	fmt.Printf("%x\n", y)  //gives the hexadecimal equivalent of y
	fmt.Printf("%#x\n", y) //gives the hexadecimal equivalent of y prefixed by 0x
	y = 911
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\t%b\t%x\n", y, y, y)
	s := fmt.Sprintf("%#x\t%b\t%x\n", y, y, y) //Sprintf is used to print a string
	fmt.Println(s)
	//fprint, fprintf, fprintln are used to print to a file
}

//OUTPUT:
// 42
// int
// 101010
// 2a
// 0x2a
// 38f
// 0x38f   1110001111      38f
// 0x38f   1110001111      38f
