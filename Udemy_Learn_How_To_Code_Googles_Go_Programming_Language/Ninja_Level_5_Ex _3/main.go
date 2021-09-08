package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	p1 := truck{
		fourWheel: true,
		vehicle:   vehicle{doors: 1, color: "Gray"},
	}

	p2 := sedan{
		luxury:  true,
		vehicle: vehicle{doors: 4, color: "Black"},
	}

	//fmt.Println(p1)
	//fmt.Println(p2)
	fmt.Printf("I have a %s truck with %v door\n", p1.color, p1.doors)
	fmt.Printf("I have a %s sedan with %v door\n", p2.color, p2.doors)
}
