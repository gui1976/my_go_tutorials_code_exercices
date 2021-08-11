package main

import (
	"fmt"
)

func main() {
	x := 33
	const y int = 36

	fmt.Printf("My number in decimal is %d\n", x)
	fmt.Printf("My number in binary is %b\n", x)
	fmt.Printf("My number in decimal is %#x\n", x)

	fmt.Printf("My number in decimal is %d\n", y)
	fmt.Printf("My number in binary is %b\n", y)
	fmt.Printf("My number in decimal is %#x\n", y)

}
