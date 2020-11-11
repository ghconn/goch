package main

import (
	"container/list"
	"fmt"
)

func main() {
	fmt.Println("za wa ru do!")
}

type cpcb func(list.List, int, list.List, int8)

var arr list.List

func compute(m int, n int, nth int, action cpcb) {
	var mixture list.List
	var index int
	var re int8
	var lpbody list.List

	loopbody(m, n, nth, index, mixture, lpbody)

	action(mixture, index, lpbody, re)
}

func loopbody(m int, n int, nth int, index int, mixture list.List, lpbody list.List) {
	if m == 0 || m == n {

		index = lpbody.Len()
	} else if m > n {
		m = m % n
		loopbody(m, n, nth, index, mixture, lpbody)
	} else {
		/* var cur=m*10%n;
		if(lpbody.indexOf(m*10)===-1){
			lpbody.push(m*10)
			m=m*10%n;
			return loopbody()
		}
		else{
			index=lpbody.indexOf(m*10)
			//lpbody=lpbody.map((el)=>(~~(el/n)))//经测试,IE浏览器不支持此es6语法
			lpbody=lpbody.map(function(el){return ~~(el/n)})
		} */
	}
}

func mapmod10(intlist list.List) {
	var temp list.List
	temp.PushBack(0)
}

/* type cb func(*[3]int32)

func compute2(cb1 cb) {
	m := [3]int32{1, 2}
	cb1(&m)
} */
