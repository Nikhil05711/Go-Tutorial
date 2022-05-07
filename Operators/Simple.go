package main

import (
	"fmt"
)

func main() {
	a := 20
	b := 33

	result1 := a + b
	fmt.Printf("The addition of a and b is %d\n", result1)

	result2 := a - b
	fmt.Printf("The sub of a and b is %d\n", result2)

	result3 := a * b
	fmt.Printf("The mul of a and b is %d\n", result3)

	result4 := a / b
	fmt.Printf("The div of a and b is %d", result4)
}
