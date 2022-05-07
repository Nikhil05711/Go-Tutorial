package main

import "fmt"

type author struct {
	name string
	age  int
}

type sponsor struct {
	salary  int
	city    string
	details author
}

func main() {
	res := sponsor{
		salary: 333, city: "Aligarh", details: author{"nikhil", 22},
	}

	fmt.Println("details of a sponsor")
	fmt.Println("sponsor salary:", res.salary)
	fmt.Println("sponsor city:", res.city)

	fmt.Println("\ndetails of a Author")
	fmt.Println("Author name:", res.details.name)
	fmt.Println("Author age:", res.details.age)

}
