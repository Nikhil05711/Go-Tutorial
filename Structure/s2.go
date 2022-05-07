package main

import "fmt"

//Access field of a structure

type car struct {
	name, color, Model string
	weight             float64
}

func main() {
	c := car{
		name: "Ferrari", color: "Black", Model: "GTAVX", weight: 2020.45,
	}

	fmt.Println("car name", c.name)
	fmt.Println("car color", c.color)

	c.color = "red"

	fmt.Println("car Details", c)
}
