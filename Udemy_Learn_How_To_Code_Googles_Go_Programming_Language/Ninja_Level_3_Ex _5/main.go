package main

import (
	"fmt"
)

func main() {
	for i := 10; i <= 100; i++ {
		fmt.Printf("Modulus of %v is %v\n", i, i%4)
	}
}
