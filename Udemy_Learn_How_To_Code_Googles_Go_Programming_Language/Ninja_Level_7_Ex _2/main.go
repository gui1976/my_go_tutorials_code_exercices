package main

import (
	"fmt"
)

type person struct {
	first_name        string
	last_name         string
	age               int
	ice_cream_flavors []string
	address           string
}

func changeMe(pp *person) {
	(*pp).address = "Palmela Village"
}

func printPerson(pp *person) {
	fmt.Printf("The ice cream flavours that %s %s likes are", pp.first_name, pp.last_name)
	for _, flavour := range pp.ice_cream_flavors {
		fmt.Printf(" %s\t", flavour)
	}
	fmt.Printf("\n")
}

func (pp person) speak() {
	fmt.Println("My name is", pp.first_name, " and I'm ", pp.age, " years old and I live in ", pp.address)
}

func main() {
	p1 := person{
		first_name:        "Gui",
		last_name:         "Lopes",
		age:               45,
		ice_cream_flavors: []string{"Ananas", "Coffee", "strawberry"},
		address:           "Lavradio",
	}

	p2 := person{
		first_name:        "Alex",
		last_name:         "Esteves",
		age:               13,
		ice_cream_flavors: []string{"Chocolate", "Mango", "Strawberry"},
		address:           "Lavradio",
	}

	m := map[string]person{
		p1.last_name: p1,
		p2.last_name: p2,
	}

	for k, v := range m {
		fmt.Printf("My key %s: ", k)
		printPerson(&v)
		v.speak()
	}

	changeMe(&p1)
	fmt.Println(p1)
}
