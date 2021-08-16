package main

import (
	"fmt"
)

func main() {
	by := 1976
	for {
		if by > 2021 {
			break
		}
		fmt.Println(by)
		by++
	}

}
