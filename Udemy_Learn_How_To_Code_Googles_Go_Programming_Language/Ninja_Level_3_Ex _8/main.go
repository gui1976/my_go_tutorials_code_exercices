package main

import (
	"fmt"
)

func main() {
	nr_day := 8

	switch {
	case nr_day == 0:
		fmt.Println("Monday")
	case nr_day == 1:
		fmt.Println("Tuesday")
	case nr_day == 2:
		fmt.Println("Wednesday")
	case nr_day == 3:
		fmt.Println("Thursday")
	case nr_day == 4:
		fmt.Println("Friday")
	case nr_day == 5:
		fmt.Println("Saturday")
	case nr_day == 6:
		fmt.Println("Sunday")
	default:
		fmt.Println("No week day")
	}
}
