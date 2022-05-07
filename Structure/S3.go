package main

import (
	"fmt"
)

//Pointers to a struct
type employee struct {
	first_name, last_name string
	pincode, age          int
}

func main() {
	e := &employee{"nikhil", "chauhan", 202001, 22}

	fmt.Println("employee firstname", (*e).first_name)
	fmt.Println("employee lastname", (*e).last_name)
	fmt.Println("employee pincode", (*e).pincode)
	fmt.Println("employee age", (*e).age)
}
