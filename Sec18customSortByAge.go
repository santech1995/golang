//Simple program to  sort data by the age given as input using the
//Sort package that implements the interface called "Interface"

package main

import (
	"fmt"
	"sort"
)

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.First, p.Age)
}

//ByAge implements sort.Interface for []Person based on
//the Age field
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

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	p1 := Person{"James", 32}
	p2 := Person{"Moneypenny", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 56}

	people := []Person{p1, p2, p3, p4}

	fmt.Println("Before sorting by age:", people)
	sort.Sort(ByAge(people))
	fmt.Println("After sorting by age:", people)
}

//OUTPUT:
// Before sorting by age: [James: 32 Moneypenny: 27 Q: 64 M: 56]
// After sorting by age: [Moneypenny: 27 James: 32 M: 56 Q: 64]
