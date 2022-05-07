package main

import "fmt"

//struct type method
type employee struct {
	name   string
	city   string
	salary int
}

func (e employee) show() {
	fmt.Println("name:", e.name)
	fmt.Println("city:", e.city)
	fmt.Println("salary:", e.salary)
}

func main() {
	res := employee{
		name:   "Nikhil",
		city:   "Aligarh",
		salary: 100,
	}

	res.show()
}
