package main

import (
	"fmt"
)

func main() {

	_my_array := []int{11, 21, 31, 41, 51, 61, 71, 81, 91, 101}

	for i, v := range _my_array {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", _my_array)
}
