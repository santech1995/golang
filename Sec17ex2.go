//Simple program to show usage of pointers with struct

package main

import (
	"fmt"
)

type person struct {
	fname string
}

func main() {
	p1 := person{
		fname: "James",
	}
	fmt.Println("Name before changeMe:", p1)
	changeMe(&p1)
	fmt.Println("Name after changeMe", p1)

}

func changeMe(p *person) {
	p.fname = "Bond" //(*p).fname <=> p.fname
}

//OUTPUT:
// Name before changeMe: {James}
// Name after changeMe {Bond}
