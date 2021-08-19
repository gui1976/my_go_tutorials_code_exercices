package main

import (
	"fmt"
)

func main() {

	_my_slice := []int{11, 21, 31, 41, 51, 61, 71, 81, 91, 101}

	_my_slice = append(_my_slice, 111)

	fmt.Println(_my_slice)

	_my_slice = append(_my_slice, 121, 131, 141)

	fmt.Println(_my_slice)

	_my_slice2 := []int{56, 57, 58, 59, 60}

	_my_slice = append(_my_slice, _my_slice2...)

	fmt.Println(_my_slice)
}
