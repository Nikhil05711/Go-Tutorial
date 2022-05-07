package main

import (
	"fmt"
)

func add(a, b int) int {
	res := a + b
	fmt.Println("result: ", res)
	return 0
}

func main() {
	fmt.Printf("It follow LIFO order\n")

	defer fmt.Println("And defer keyword play when all method runs")
	defer add(10, 45)
	defer add(13, 45)
}
