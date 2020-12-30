package main

import (
	"fmt"
)

func main() {
	//func (r reciever identifier(parameters) return(s)) {...}
	foo()
	bar("James")
	s1 := woo("Mary")
	fmt.Println(s1)
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x)
	fmt.Println(y)
}

func foo() {
	fmt.Println("hello from foo")
}

//func taking a parameter:
func bar(s string) {
	fmt.Println("Hello ", s)
}

//func returning a string:
func woo(st string) string {
	return fmt.Sprint("Hello from woo ", st)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `,says "Hello`)
	b := true
	return a, b
}

//we declare functions with 'parameters' and
//we call functions by passing 'arguments'
//everything in Go is PASS BY VALUE
