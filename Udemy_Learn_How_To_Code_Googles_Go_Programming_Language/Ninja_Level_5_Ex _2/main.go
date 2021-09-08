package main

import (
	"fmt"
)

type person struct {
	first_name        string
	last_name         string
	ice_cream_flavors []string
}

func printPerson(pp *person) {
	fmt.Printf("The ice cream flavours that %s %s likes are", pp.first_name, pp.last_name)
	for _, flavour := range pp.ice_cream_flavors {
		fmt.Printf(" %s\t", flavour)
	}
	fmt.Printf("\n")
}

func main() {
	p1 := person{
		first_name:        "Gui",
		last_name:         "Lopes",
		ice_cream_flavors: []string{"Ananas", "Coffee", "strawberry"},
	}

	p2 := person{
		first_name:        "Alex",
		last_name:         "Esteves",
		ice_cream_flavors: []string{"Chocolate", "Mango", "Strawberry"},
	}

	m := map[string]person{
		p1.last_name: p1,
		p2.last_name: p2,
	}

	for k, v := range m {
		fmt.Printf("My key %s: ", k)
		printPerson(&v)
	}
}
