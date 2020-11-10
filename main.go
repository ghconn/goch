package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 猜数字游戏，定义一个要猜的数字，由rand.Intn函数产生随机数与定义的数进行比较
func main() {
	num := 996  // 要猜的数字
	min := 0    // 产生随机数区间最小值
	max := 1001 // 产生随机数区间最大值
	rand.Seed(time.Now().UnixNano())
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
	fmt.Printf("n:%v, greet!\n", n)
}
