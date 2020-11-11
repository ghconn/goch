package main

import (
	"container/list"
	"fmt"
)

func main() {
	m := 22
	n := 122
	nth := 100
	compute(m, n, nth, func(mixture *list.List, index int, lpbody *list.List, re int8) {
		if mixture.Len() > 0 {
			fmt.Println("mixture:", mixture, "从第", index, "位出现循环体")
		} else {
			fmt.Println("从第1位出现循环体")
		}
		fmt.Println()
		fmt.Printf("小数部分第%v位的数字:%v\n", n, re)
		fmt.Println()
		fmt.Println("循环体：")
		fmt.Println(lpbody)
	})
}

type cpcb func(*list.List, int, *list.List, int8)

var arr list.List

func compute(m int, n int, nth int, action cpcb) { // nth 要求的第几位
	mixture := list.New() // mixture 计算出的混合体
	var index int
	var re int8
	lpbody := list.New()

	//m 被除数 n 除数 index 计算出的第几位出现循环体 lpbody 计算出的循环体
	loopbody(m, n, index, lpbody)

	if index > 0 {
		mixture = slice(lpbody, 0, index)
	}

	nth = nth - 1
	if nth < index {
		re = elementAt(lpbody, nth).(int8)
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
			re = int8(indexOf(lpbody, (nth-index)%lpbody.Len()))
		}
	}

	action(mixture, index, lpbody, re)
}

func loopbody(m int, n int, index int, lpbody *list.List) {
	if m == 0 || m == n {
		mapmod10(lpbody)
		index = lpbody.Len()
	} else if m > n {
		m = m % n
		loopbody(m, n, index, lpbody)
	} else {
		if indexOf(lpbody, m*10) == -1 {
			lpbody.PushBack(m * 10)
			m = m * 10 % n
			loopbody(m, n, index, lpbody)
		} else {
			index = indexOf(lpbody, m*10)
			mapdiv(lpbody, n)
		}
	}
}

func mapmod10(intlist *list.List) {
	temp := list.New()
	for x := intlist.Front(); x != nil; x = x.Next() {
		temp.PushBack(x.Value.(int) % 10)
	}
	intlist = temp
}

func mapdiv(intlist *list.List, divisor int) {
	temp := list.New()
	for x := intlist.Front(); x != nil; x = x.Next() {
		temp.PushBack(x.Value.(int) / divisor)
	}
	intlist = temp
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

func slice(lst *list.List, start int, end int) *list.List { // [start, end]
	var temp *list.List
	for i, x := 0, lst.Front(); x != nil && i <= end; i, x = i+1, x.Next() {
		if i >= start {
			temp.PushBack(x.Value)
		}
	}
	return temp
}

func splice(lst *list.List, start int, end int) { // [start, end]
	var tempelement *list.Element
	for i, x := 0, lst.Front(); x != nil && i <= end; i, x = i+1, tempelement {
		if i < start || i > end {
			tempelement = x.Next()
			lst.Remove(x)
		}
	}
}
