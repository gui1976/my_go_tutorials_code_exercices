package main

import (
	"fmt"
)

func foo() int {
	return 1976
}

func bar() (string, int) {
	return "Gui", 45
}

func main() {

	name, shirt_number := bar()
	fmt.Println("My name is ", name, ", I've born in ", foo(), "and my baseball jersey number is ", shirt_number)
}
