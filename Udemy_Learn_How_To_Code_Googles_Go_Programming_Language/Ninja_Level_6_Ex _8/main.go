package main

import "fmt"

func yet_a_func() func() int {
	return func() int {
		return 1976
	}
}

func main() {
	x := yet_a_func()
	fmt.Println(x())
}
