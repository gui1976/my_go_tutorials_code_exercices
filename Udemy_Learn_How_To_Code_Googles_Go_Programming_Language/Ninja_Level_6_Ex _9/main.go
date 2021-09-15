package main

import "fmt"

func yet_a_func() func() int {
	return func() int {
		return 1976
	}
}

func printSomeStuff(f func() int) {
	fmt.Println("I've have born in the year ", f())
}

func main() {
	printSomeStuff(yet_a_func())
}
