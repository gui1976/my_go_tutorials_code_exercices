package main

import (
	"fmt"
)

func main() {

	_my_array := [5]int{1, 2, 3, 4, 5}

	for i, v := range _my_array {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", _my_array)
}
