package main

import (
	"fmt"
)

func main() {

	_my_slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	_my_slice = append(_my_slice[:3], _my_slice[6:]...)

	fmt.Println(_my_slice)
}
