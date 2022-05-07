package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

func main() {

	//Sorting the given slice
	a := []int{45, 65, 24, 54, 76, 67, 84}
	sort.Ints(a)
	fmt.Printf("The sorted slice is: %d\n", a)

	//finding the particular index in an string
	fmt.Println("Index value is:", strings.Index("GeeksForGeeks\n", "ks"))

	//Find the current time in golang
	fmt.Println("Time:", time.Now().Unix())
}
