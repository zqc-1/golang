package main

import "fmt"

/*
函数功能：实现代码模块化，避免重复

函数声明

	func 函数名（）{
		功能代码
	}

函数调用：
函数名（）
*/
func printLing() {
	// 打印菱形
	// 层数
	var n int = 6

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
func main() {

	//打印菱形
	printLing()
	fmt.Println("==============================")
	printLing()
}
