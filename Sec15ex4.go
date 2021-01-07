//Simple program to illustrate the use of methods in GO

package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

//func receiver identifier parenthesis return

func (p person) speak() {
	fmt.Println("My name is: ", p.first, p.last)
	fmt.Println("My age is: ", p.age)
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   34,
	}
	p1.speak()

}

//OUTPUT:
// My name is:  James Bond
// My age is:  34
