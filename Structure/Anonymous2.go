package main

//Anonymous fields
import "fmt"

type author struct {
	string
	int
}

func main() {
	a := author{"nikhil", 22}

	fmt.Println(a)
}
