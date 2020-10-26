package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	num := 996
	min := 0
	max := 1001
	n := rand.Intn(max-min-1) + min + 1
	for num != n {
		fmt.Printf("数字为：%2v\n", n)
		if n > num {
			max = n
		} else {
			min = n
		}
		n = rand.Intn(max-min-1) + min + 1
		time.Sleep(time.Second)
	}
	fmt.Println("greet!")
}
