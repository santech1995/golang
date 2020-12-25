//Program to implement struct data structure in GO

//Struct is a data structure which allows us to compose together
//values of different data types

package main

import (
	"fmt"
)

//type struct_name struct {       //struct_name is called 'type'
//field1 type_of_field1,
//field2 type_of_field2,
//}
//object1 := struct_name {         //object1 is called 'value'
//	field1: value1,
//	field2: value2,
//}
//accessing the fields inside struct using dot operator
//object1.field1

type person struct {
	first string
	last  string
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
	}
	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
	}
	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.first, p1.last)
	fmt.Println(p2.first, p2.last)
}

//OUTPUT:-
// {James Bond}
// {Miss Moneypenny}
// James Bond
// Miss Moneypenny
