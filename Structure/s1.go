package main

import (
	"fmt"
)

type Address struct {
	Name    string
	city    string
	pincode int
}

func main() {
	var a Address
	fmt.Println(a)

	a1 := Address{"Aniket", "Lucknow", 206001}
	fmt.Println(a1)

}
