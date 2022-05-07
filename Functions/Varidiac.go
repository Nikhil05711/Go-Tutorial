package main

import (
	"fmt"
	"strings"
)

func joinStr(element ...string) string {
	return strings.Join(element, "")
}

func main() {
	fmt.Println(joinStr())
	fmt.Println(joinStr("geeks", "for", "geeks"))
}
