package main

import (
	"fmt"
)

type Author struct {
	name string
	age  int
}

type HR struct {
	details Author
}

func main() {
	result := HR{
		details: Author{"Nikhil", 22},
	}

	fmt.Println("Details of the author")
	fmt.Println(result)
}
