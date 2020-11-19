//Simple program to study the type of data in GO

package main

import "fmt"

var y = 43

//Declare the variable with the identifier "z" as being of type string
//and assign it to "Shaken,not stirred"
//GO is a static programming lang ie once "z" is assigned a string, it cannot be assigned data of
//any other type
//like in java viz a dynamic lang
var z = "Shaken,not stirred"

func main() {
	fmt.Println(y)
	fmt.Printf("y is an %T\n", y)
	//fmt.Println(z)
	//z = "happy"
	fmt.Printf("z is a %T\n", z)
	fmt.Println(z)
	//z = 43 is not allowed because z is already assigned a string and cannot be assigned an int
}

//Output:
// 43
// y is an int
// z is a string
// Shaken,not stirred
