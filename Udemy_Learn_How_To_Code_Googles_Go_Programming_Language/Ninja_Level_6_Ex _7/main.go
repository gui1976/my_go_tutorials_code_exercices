package main

import (
	"fmt"
)

func main() {
	x := func() {
		fmt.Println("Just a func assigned to a variable")
	}

	x()
}
