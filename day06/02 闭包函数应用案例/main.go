package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

func foo() {

	fmt.Println("foo start")
	time.Sleep(time.Second * 2)
	fmt.Println("foo end")

}

func bar() {
	fmt.Println("bar start")
	time.Sleep(time.Second * 1)
	fmt.Println("bar end")
}

func getCalledNum(f func()) func() { //计数器装饰函数
	var count = 0

	calledNum := func() {

		f()
		count++
		funcName := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
		fmt.Printf("%s函数被调用地第 %d 次\n", funcName, count)
	}
	return calledNum
}

func getProTimer(f func()) func() { //计数器装饰函数

	return func() {

		start := time.Now().Unix()
		f()
		end := time.Now().Unix()
		fmt.Println("cost time:", end-start)
	}

}
func main() {

	/*calledNumFoo := getCalledNum(foo)
	calledNumFoo()
	calledNumFoo()
	calledNumFoo()

	calledNameBar := getCalledNum(bar)
	calledNameBar()
	calledNameBar()
	calledNameBar()*/

	/*foo := getCalledNum(foo)
	foo()
	foo()
	foo()*/

	foo := getProTimer(foo)
	foo()

	bar := getProTimer(bar)
	bar()
}
