package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var emoji rune = 128515
	fmt.Printf("%c\n", emoji)
	emoje2 := "😃"
	fmt.Println(emoje2)
	i := 65
	fmt.Printf("%c\n", i)

	emoje3 := "😃xyz"
	fmt.Println(len(emoje3))
	fmt.Println(utf8.RuneCountInString(emoje3))
}
