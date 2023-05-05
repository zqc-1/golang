package main

import "fmt"

var x = 100 //全局变量 不能声明成 下：= 100

func foo() {
	var x = 1
	fmt.Println(x)
	fmt.Println("foo...")
}

func bar() {
	foo()
	fmt.Println("bar...")
	fmt.Println(x)

}

func main() {
	bar()
}
