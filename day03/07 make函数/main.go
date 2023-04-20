package main

import "fmt"

func main() {
	//初始化创建空间
	var s = make([]int, 5, 10)
	fmt.Println(len(s), cap(s))
	fmt.Println(s)
	s[0] = 100
	fmt.Println(s)

	a := make([]int, 5)
	b := a[0:3]
	a[0] = 100
	fmt.Println(a)
	fmt.Println(b)
}
