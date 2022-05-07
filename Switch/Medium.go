package main

import (
	"fmt"
)

var value int = 5

func main() {
	switch {
	case value == 1:
		fmt.Println("Sunday")

	case value == 2:
		fmt.Println("Monday")

	case value == 3:
		fmt.Println("Tuesday")

	case value == 4:
		fmt.Println("Wednesday")

	case value == 5:
		fmt.Println("Thursday")

	case value == 6:
		fmt.Println("Friday")

	case value == 7:
		fmt.Println("Saturday")

	default:
		fmt.Println("Invalid")
	}
}
