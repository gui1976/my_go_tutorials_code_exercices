package main

import (
	"fmt"
)

func main() {

	favSport := "Baseball"

	switch favSport {
	case "Baseball":
		fmt.Println("Play ball.")
	case "Basket":
		fmt.Println("It's part of the game.")
	case "Football":
		fmt.Println("Penalty!!!")
	}
}
