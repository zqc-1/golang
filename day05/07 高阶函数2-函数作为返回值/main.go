package main

import (
	"fmt"
)

func foo() func() int {
	var inner = func() int {
		fmt.Println("一个新的返回值")
		return 100
	}
	//fmt.Println(reflect.TypeOf(inner))
	return inner
}

func main() {
	//var f func() int
	//f = foo() //返回inner函数赋值给f变量
	//f()       //函数调用

	foo()()
}
