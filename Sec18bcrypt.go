//Program to implement the bcrypt concept to generate password and its hash

package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Password before encryption is:", s)
	fmt.Println("Password after encryption is(hashed version) :", bs)

	loginPword1 := `password123`

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPword1))
	if err != nil {
		fmt.Println("YOU CAN'T LOGIN")
	}

	fmt.Println("You are logged in")

}

//OUTPUT:
// Password before encryption is: password123
// Password after encryption is(hashed version) : [36 50 97 36 48 52 36
//  87 81 74 67 110 86 122 107 70 87 107 56 90 85 47 114 76 67 117 51 114
//   117 68 65 85 51 99 49 114 88 118 82 112 51 97 73 75 79 53 67 117 90
//    48 99 90 97 65 83 50 104 121 111 121]

//You are logged in
