package main

import (
	"fmt"
)

func main() {
	var s string = "shalom"

	l := len(s)
	var i int
	for i < l {
		fmt.Printf("%c\n", s[i])
		i++
	}
}
