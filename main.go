package main

import (
	"container/list"
	"fmt"
)

func main() {
	m := 229
	n := 1222
	nth := 2000
	compute(m, n, nth, func(mixture *list.List, lpbody *list.List, re int8) {
		if mixture.Len() > 0 {
			fmt.Print("mixture:")
			loopthroughprint(mixture)
			fmt.Printf("从第%v位出现循环体", index+1)
		} else {
			fmt.Println("从第1位出现循环体")
		}
		fmt.Println()
		fmt.Printf("小数部分第%v位的数字:%v\n", nth, re)
		fmt.Println()
		fmt.Println("循环体：")
		loopthroughprint(lpbody)
		fmt.Println("循环体长度：", lpbody.Len())
	})
}

type cpcb func(*list.List, *list.List, int8)

var index int

func compute(m int, n int, nth int, action cpcb) { // nth 要求的第几位
	mixture := list.New() // mixture 计算出的混合体
	var re int8
	lpbody := list.New()

	//m 被除数 n 除数 index 计算出的第几位出现循环体 lpbody 计算出的循环体
	loopbody(m, n, lpbody)

	if index > 0 {
		slicetoanother(lpbody, mixture, 0, index)
	}

	nth = nth - 1
	if nth < index {
		re = int8(elementAt(lpbody, nth).(int))
		if index == lpbody.Len() {
			lpbody = list.New()
			lpbody.PushBack(0)
		} else {
			splice(lpbody, 0, index)
		}
	} else {
		if index == lpbody.Len() {
			lpbody = list.New()
			lpbody.PushBack(0)
		} else {
			splice(lpbody, 0, index)
			re = int8(elementAt(lpbody, (nth-index)%lpbody.Len()).(int))
		}
	}

	action(mixture, lpbody, re)
}

func loopbody(m int, n int, lpbody *list.List) {
	if m == 0 || m == n {
		mapdiv(lpbody, n)
		index = lpbody.Len()
	} else if m > n {
		m = m % n
		loopbody(m, n, lpbody)
	} else {
		index = indexOf(lpbody, m*10)
		if index == -1 {
			lpbody.PushBack(m * 10)
			m = m * 10 % n
			loopbody(m, n, lpbody)
		} else {
			mapdiv(lpbody, n)
		}
	}
}

func mapmod10(intlist *list.List) {
	for x := intlist.Front(); x != nil; x = x.Next() {
		x.Value = x.Value.(int) % 10
	}
}

func mapdiv(intlist *list.List, divisor int) {
	for x := intlist.Front(); x != nil; x = x.Next() {
		x.Value = x.Value.(int) / divisor
	}
}

func indexOf(lst *list.List, v interface{}) int {
	for i, x := 0, lst.Front(); x != nil; x, i = x.Next(), i+1 {
		if x.Value.(int) == v.(int) {
			return i
		}
	}
	return -1
}

func elementAt(lst *list.List, index int) interface{} {
	for i, x := 0, lst.Front(); x != nil; x, i = x.Next(), i+1 {
		if index == i {
			return x.Value
		}
	}
	return nil
}

func slicetoanother(lst *list.List, otherlst *list.List, start int, end int) { // [start, end)
	for i, x := 0, lst.Front(); x != nil && i < end; i, x = i+1, x.Next() {
		if i >= start {
			otherlst.PushBack(x.Value)
		}
	}
}

func splice(lst *list.List, start int, end int) { // [start, end)
	var tempelement *list.Element
	for i, x := 0, lst.Front(); x != nil && i < end; i, x = i+1, tempelement {
		if i >= start {
			tempelement = x.Next()
			lst.Remove(x)
		}
	}
}

// 遍历打印
func loopthroughprint(lst *list.List) {
	for x := lst.Front(); x != nil; x = x.Next() {
		fmt.Print(x.Value)
		if x.Next() != nil {
			fmt.Print(",")
		}
	}
	fmt.Println()
}
