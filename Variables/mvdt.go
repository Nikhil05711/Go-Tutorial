package main

import (
	"fmt"
)

func main() {
	var value1, value2, value3 = 22, "March", 19.99
	fmt.Printf("The value of value1 is %d\n", value1)
	fmt.Printf("The Type of value1 is %T\n", value1)

	fmt.Printf("The value of value2 is %s\n", value2)
	fmt.Printf("The Type of value2 is %T\n", value2)

	fmt.Printf("The value of value3 is %f\n", value3)
	fmt.Printf("The Type of value3 is %T", value3)

}
