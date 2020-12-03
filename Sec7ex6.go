//Program that creates 4 constants for the last 4 years using "iota"

package main

import "fmt"

const (
	a = 2016 + iota
	b = 2016 + iota
	c = 2016 + iota
	d = 2016 + iota
)

func main() {
	fmt.Println(a, b, c, d)
}

//OUTPUT:
// 2016 2017 2018 2019
