package main

import (
	"fmt"
)

func main() {

	//_007 := make([][]string, 3)

	_007 := [][]string{
		[]string{"James", "Bond", "Shaken, not stirred"},
		[]string{"Miss", "Moneypenny", "Helloo, James."},
	}

	fmt.Println(_007)

	for i, xs := range _007 {

		fmt.Println("record", i)

		for y, ys := range xs {
			fmt.Printf("Index position: %v \t value: %s\n", y, ys)
		}
	}
}
