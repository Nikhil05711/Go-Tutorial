package main

import (
	"fmt"
)

func main() {
	var i int = 25

	if i < 100 {
		fmt.Printf("The value of i is smaller then 100: %d", i)
	}

	fmt.Printf("\nThe value of i is %d", i)
}
