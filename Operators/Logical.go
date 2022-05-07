package main

import (
	"fmt"
)

func main() {
	var p int = 26
	var q int = 54

	if p != q && p <= q {
		fmt.Println("true")
	}

	if p != q || p >= q {
		fmt.Println("true")
	}

	if !(p == q) {
		fmt.Println("true")
	}
}
