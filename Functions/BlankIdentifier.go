package main

import "fmt"

func main() {

	mul, _ := mul_div(100, 100)
	fmt.Println("100 * 100\n", mul)
}

func mul_div(n1 int, n2 int) (int, int) {
	return n1 * n2, n1 / n2
}
