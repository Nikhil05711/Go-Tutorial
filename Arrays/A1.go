package main

import "fmt"

func main() {
	var myint [2]int

	myint[0] = 1
	myint[1] = 100

	fmt.Println("Details of my integers")
	fmt.Println("myint 0:", myint[0])
	fmt.Println("myint 1:", myint[1])

}
