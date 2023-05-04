package main

import "fmt"

func foo() {
	//var x int
	//x = 10

	x = 20
	fmt.Println(x)
	fmt.Println(&x)

}

func bar() {
	var x = 100
	x = 200
	fmt.Println(x)
}

// 全局变量
var x = 1

func main() {

	foo()
	//var x = 100
	fmt.Println(x)
	fmt.Println(&x)

	// if for 都可以开辟作用域

	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

}
