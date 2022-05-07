package main

import (
	"fmt"
)

func main() {
	var i int = 65

	if i < 100 {
		fmt.Printf("The value of i is less then 100: %d", i)
	} else {
		fmt.Printf("The value is: %d", i)
	}
}
