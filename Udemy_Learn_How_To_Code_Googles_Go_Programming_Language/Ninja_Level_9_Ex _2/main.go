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

type human interface {
	speak()
}

func (pp *person) speak() {
	fmt.Println("My name is", pp.first_name, "and I'm", pp.age, "years old and I live in", pp.address)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first_name:        "Gui",
		last_name:         "Lopes",
		age:               45,
		ice_cream_flavors: []string{"Ananas", "Coffee", "strawberry"},
		address:           "Quinta do Anjo",
	}

	saySomething(&p1)
}
