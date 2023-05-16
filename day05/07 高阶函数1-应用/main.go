package main

import (
	"fmt"
	"time"
)

func foo() {
	start := time.Now().Unix()
	fmt.Println("foo功能开始")
	time.Sleep(time.Second * 2)
	fmt.Println("foo功能结束")
	end := time.Now().Unix()
	fmt.Println("cost time:", end-start)
}
func bar() {
	fmt.Println("bar功能开始")
	time.Sleep(time.Second * 3)
	fmt.Println("bar功能结束")
}
func timer(f func()) {
	start := time.Now().Unix()
	f()
	end := time.Now().Unix()
	fmt.Println("timer:", end-start)
}
func main() {

	/*f1 := time.Now().Unix()
	fmt.Println(f1, reflect.TypeOf(f1)) //1684216035 int64

	time.Sleep(time.Second * 3)

	f2 := time.Now().Unix()
	fmt.Println(f2, reflect.TypeOf(f2))*/
	//foo()
	//bar()

	timer(foo)
	timer(bar)
}
