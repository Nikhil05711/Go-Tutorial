package main

import (
	"fmt"
)

func main() {
	mmap := map[int]string{
		11: "First",
		22: "second",
		33: "third",
	}
	for key, value := range mmap {
		fmt.Println(key, value)
	}
}
