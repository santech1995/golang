package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}
	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}
	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)
	b, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	os.Stdout.Write(b)
}

//OUTPUT:
// [{James 32} {Moneypenny 27} {M 54}]
// [{"First":"James","Age":32},{"First":"Moneypenny","Age":27},{"First":"M","Age":54}]
