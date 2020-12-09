//Simple program to implement maps in GO

package main

import (
	"fmt"
)

func main() {
	//map_name := map[key_type]value_type{
	//key1: value1,
	//key2: value2,
	//}

	m := map[string]int{
		"James": 32,
		"Mary":  21,
	}
	fmt.Println(m)
	fmt.Println(m["James"]) //value = 32
	fmt.Println(m["Judy"])  //Since this key does not exist its value = 0(default)
	//comma-ok idiom : used to check if a value is present in the map
	v, ok := m["Judy"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Judy"]; ok {
		fmt.Println("This is the IF print", v)
	}
	//the above syntax is euivalent to ' v,ok := m["Judy"] '.It will
	//be printed only if "Judy" exists in the map as key
}

//OUTPUT:
// map[James:32 Mary:21]
// 32
// 0
// 0
// false
