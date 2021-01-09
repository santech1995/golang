//Simple program to calculate area of square and circle using
//structs and interfaces

package main

import (
	"fmt"
	"math"
)

type square struct {
	side float32
	// area int
}

func (s square) area() float32 {
	return s.side * s.side
}

type circle struct {
	radius float32
	// area float32
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float32
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	s1 := square{
		side: 10,
	}
	c1 := circle{
		radius: 10,
	}
	info(s1)
	info(c1)
}

//OUTPUT:

// 100
// 314.15927
