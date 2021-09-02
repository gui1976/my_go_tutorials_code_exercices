package main

import (
	"fmt"
)

func main() {

	m := map[string][]string{
		`Bond_James`:      []string{`Shaken`, `not stirred`, `women`},
		`Moneypenny_Miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`No_Dr`:           []string{`Being Evil`, `Ice cream`, `Sunsets`},
	}

	// fmt.Println(m)

	for k, v := range m {

		fmt.Println("This is the record for ", k)
		for i, v2 := range v {
			fmt.Println("\t ", i, v2)
		}
	}
}
