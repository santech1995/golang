//Program to implement interfaces in GO

package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, "--secretAgent Speak()")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " --person Speak()")
}

type human interface {
	speak()
}

//sa1 value is of type secretAgent and human bacause it has access to
//the speak() method

//even the p1 value is of type secretAgent and human becoz it
//has access to speak() method

//an interface allows a value to be of different types

func bar(h human) {
	fmt.Println("I was passed into bar", h)
}

func main() {
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}
	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}

	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}
	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)

	bar(sa1)
	bar(sa2)
	bar(p1)
	//function bar can take in values of many types => polymorphism
}

//OUTPUT:

// {{James Bond} true}
// I am James Bond --secretAgent Speak()
// I am Miss Moneypenny --secretAgent Speak()
// {Dr. Yes}
// I was passed into bar {{James Bond} true}
// I was passed into bar {{Miss Moneypenny} true}
// I was passed into bar {Dr. Yes}
