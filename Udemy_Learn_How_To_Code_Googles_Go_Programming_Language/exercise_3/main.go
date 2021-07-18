package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {

	result := fmt.Sprintf("It's %t that %s has %d years old.\n", z, y, x)

	fmt.Println(result)

}
