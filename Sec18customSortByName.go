//Simple program to sort a slice by the name given as input
//using the sort package that implements the interface called
//"Interface"

package main

import (
	"fmt"
	"sort"
)

//ByName implements sort.Interface for []Person based on
//the First field
// type Interface interface {
// 	     //Len is the number of elemetns in the collection
// 	     Len() int
// 	     //Less reports whether the element with index i should sort
// 	     //before the element with index j.
// 	     Less(i,j int) bool
// 	     //Swap swaps the elements with indexes i and j.
// 	     Swap(i,j int)
// }

type Person struct {
	First string
	Age   int
}

type ByName []Person

func (bn ByName) Len() int           { return len(bn) }
func (bn ByName) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
func (bn ByName) Less(i, j int) bool { return bn[i].First < bn[j].First }

func main() {
	p1 := Person{"James", 32}
	p2 := Person{"Moneypenny", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 56}

	people := []Person{p1, p2, p3, p4}

	fmt.Println("Before sorting by name:", people)
	sort.Sort(ByName(people))
	fmt.Println("After sorting by name:", people)
}

//OUTPUT:
// Before sorting by name: [{James 32} {Moneypenny 27} {Q 64} {M 56}]
// After sorting by name: [{James 32} {M 56} {Moneypenny 27} {Q 64}]
