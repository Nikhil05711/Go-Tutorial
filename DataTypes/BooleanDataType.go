package main

import (
	"fmt"
)

func main() {
	string1 := "Nikhil_chauhan"
	string2 := "Aniket_jaiswal"
	string3 := "Nikhil_chauhan"

	result1 := string1 == string2
	result2 := string1 == string3

	fmt.Println(result1)
	fmt.Println(result2)

	fmt.Printf("The type of result1 is %T and "+"The type of result2 is %T", result1, result2)
}
