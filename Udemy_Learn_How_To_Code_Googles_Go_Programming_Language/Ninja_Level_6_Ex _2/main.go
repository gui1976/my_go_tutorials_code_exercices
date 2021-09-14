package main

import (
	"fmt"
)

func foo(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}

	return total
}

func bar(x []int) int {
	total := 0
	for _, v := range x {
		total += v + 1
	}

	return total
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(foo(numbers...))
	fmt.Println(bar(numbers))
}
