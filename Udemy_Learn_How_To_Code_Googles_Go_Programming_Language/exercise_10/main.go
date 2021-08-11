package main

import (
	"fmt"
)

const (
	year_1 = iota + 2018
	year_2
	year_3
	year_4
)

func main() {

	fmt.Printf("%d\t%d\t%d\t%d\t", year_1, year_2, year_3, year_4)

}
