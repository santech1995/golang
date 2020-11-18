/*Simple program to convert a variable of custom datatype(hotdog) into a
  built-in(int) type*/

package main

import "fmt"

var a int

type hotdog int

var b hotdog

func main() {
	a = 42
	fmt.Println("a = ", a, "\n")
	fmt.Printf("a is of type %T\n", a)
	b = 30
	fmt.Println("b=", b, "\n")
	fmt.Printf("b is of type %T\n", b)
	a = int(b) //conversion of a value of type 'hotdog' to type 'int'
	fmt.Println(a)
	fmt.Printf("a is of type %T\n", a)
	fmt.Printf("b is of type %T\n", b)
}

/*
OUTPUT:
a =  42

a is of type int
b= 30

b is of type main.hotdog
30
a is of type int
b is of type main.hotdog */
