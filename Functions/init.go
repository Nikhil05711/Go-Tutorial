package main

import (
	"fmt"
)

func init() {
	fmt.Print("This is the init function\n")
}

func init() {
	fmt.Print("Multiple init function i had used here\n")
}

func main() {
	fmt.Println("main function always come after init function")
}
