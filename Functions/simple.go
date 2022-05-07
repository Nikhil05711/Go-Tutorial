package main

import (
	"fmt"
)

func area(length, width int) int {
	ar := length * width
	return ar
}

func main() {
	fmt.Printf("The area is: %d", area(10, 49))
}
