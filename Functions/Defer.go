package main

import (
	"fmt"
)

func mul(a, b int) int {
	res := a * b
	fmt.Printf("The result is %d\n", res)
	return 0
}

func show() {
	fmt.Println("Geeks for geeks")
}

func main() {
	mul(10, 34)
	defer mul(10, 32)
	show()
}
