package main

import (
	"fmt"
	"reflect"
)

func printLing(n int) {
	// 打印菱形
	// 层数

	for i := 1; i <= n; i++ {
		for k := 1; k <= n-i; k++ {
			fmt.Print(" ")
		}

		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := n - 1; i >= 1; i-- {
		for k := 1; k <= n-i; k++ {
			fmt.Print(" ")
		}

		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func sum(x, y int) {
	var res = 0
	for i := x; i <= y; i++ {
		res += i
	}
	fmt.Println(res)
}

// 不限个护士加法器
func add(s ...int) {
	fmt.Println(s, reflect.TypeOf(s))
	var res = 0
	for _, v := range s {
		res += v
	}
	fmt.Println(res)
}

func main() {

	//printLing(4)
	//sum(1, 100)
	//sum(50, 1050)
	add(1, 2, 3, 4)
}
