package main

import (
	"fmt"
	"reflect"
)

func foo() {
	var bar func(int, int) int
	bar = func(x, y int) int {
		return 100
	}
	ret := bar(1, 1)
	fmt.Println(ret, reflect.TypeOf(bar))
}

func main() {

	var f = func() {
		fmt.Println("hello world!!!")
	}

	f()

	(func() {
		fmt.Println("hello world!!!")
	})()

	(func(x, y int) {
		fmt.Println(x + y)
	})(10, 15)

	foo()
}
