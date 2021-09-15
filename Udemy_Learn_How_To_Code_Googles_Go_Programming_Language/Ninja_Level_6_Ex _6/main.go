package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("Just an anonymous func")
	}()
}
