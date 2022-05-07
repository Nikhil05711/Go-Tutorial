package main

import (
	"fmt"
)

func main() {
	var a complex128 = complex(4, 5)
	var b complex64 = complex(5, 3)

	fmt.Println(a)
	fmt.Println(b)

	fmt.Printf("The type of %T and"+" The type of %T ", a, b)
}
