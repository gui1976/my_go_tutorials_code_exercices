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
		First: "Gui",
		Age:   45,
	}

	u2 := user{
		First: "Xano",
		Age:   13,
	}

	u3 := user{
		First: "André",
		Age:   19,
	}

	users := []user{u1, u2, u3}

	//fmt.Println(users)

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println("error: ", err)
	}

	os.Stdout.Write(bs)
}
