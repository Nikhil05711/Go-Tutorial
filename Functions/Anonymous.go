package main

import (
	"fmt"
)

func main() {
	//Anonymous function
	func() {
		fmt.Printf("Welcome to anonymous function\n")
	}()

	//when function assign to a variable
	value := func() {
		fmt.Printf("welcome again\n")
	}
	value()

	//when we pass arguments
	func(ele string) {
		fmt.Printf(ele)
	}("once again we welcome you")

}
