//Simple program to illustrate encoding of data into json format

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken,not stirred",
			"Youth is no gaurantee of innovation",
			"In his majesty's royal service",
		},
	}
	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James,it is soo good to see you",
			"Would you like to take care of that for you,James?",
			"I would really prefer to be a secret agent myself",
		},
	}
	u3 := user{
		First: "M",
		Last:  "Hmmmmm",
		Age:   54,
		Sayings: []string{
			"Oh,James.You didn't.",
			"Oh God,what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}
	users := []user{u1, u2, u3}
	fmt.Println(users)

	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println("We did something wrong and here's the error:", err)
	}

}

//OUTPUT:
// [{James Bond 32 [Shaken,not stirred Youth is no gaurantee of
//innovation In his majesty's royal service]}
//{Miss Moneypenny 27 [James,it is soo good to see you Would you like
// to take care of that for you,James? I would really prefer to be a
// secret agent myself]}
//{M Hmmmmm 54 [Oh,James.You didn't. Oh God,what has James done now?
//Can someone please tell me where James Bond is?]}]

// [{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken,not
// stirred","Youth is no gaurantee of innovation","In his majesty's
// royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,
//"Sayings":["James,it is soo good to see you","Would you like to take
// care of that for you,James?","I would really prefer to be a secret
// agent myself"]},{"First":"M","Last":"Hmmmmm","Age":54,"Sayings":
//["Oh,James.You didn't.","Oh God,what has James done now?",
//"Can someone please tell me where James Bond is?"]}]
