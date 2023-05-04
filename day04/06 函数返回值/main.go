package main

import (
	"fmt"
	"reflect"
)

// 无返回值
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

func add(a, b int) int {
	c := a + b
	return c //函数终止语句
}

// 返回值命名
func add2(s ...int) (z int) {
	fmt.Println(s, reflect.TypeOf(s))
	var res = 0
	for _, v := range s {
		res += v
	}
	return
}

func login(user, pwd string) (string, string) {
	if user == "root" && pwd == "root" {
		return "登录成功！", user
	} else {
		return "账户名或密码错误！", ""
	}
}

func main() {
	res := add(1, 2)
	fmt.Println(res)
	add(res, 100)

	msg, user := login("root", "root")
	fmt.Println(msg, user)

	//返回值命名
	fmt.Println(add2(1, 2, 3))
}
