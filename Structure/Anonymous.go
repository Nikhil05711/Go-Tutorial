package main

//Anonymous struct
import (
	"fmt"
)

func main() {
	e := struct {
		name string
		age  int
	}{
		name: "nikhil",
		age:  22,
	}

	fmt.Println(e)
}
