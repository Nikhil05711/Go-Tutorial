package main

import "fmt"

func main() {
	a1 := [3]string{"aa", "vv", "cc"}
	a2 := a1
	a1[0] = "ss"

	fmt.Println(a1)
	fmt.Println(a2)
}
