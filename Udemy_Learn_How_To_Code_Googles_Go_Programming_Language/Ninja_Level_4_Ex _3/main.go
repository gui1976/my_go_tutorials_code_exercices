package main

import (
	"fmt"
)

func main() {

	_my_array := []int{11, 21, 31, 41, 51, 61, 71, 81, 91, 101}

	slice_1 := _my_array[:5]
	slice_2 := _my_array[5:10]
	slice_3 := _my_array[2:7]
	slice_4 := _my_array[1:6]

	fmt.Println(_my_array)
	fmt.Println(slice_1)
	fmt.Println(slice_2)
	fmt.Println(slice_3)
	fmt.Println(slice_4)
}
