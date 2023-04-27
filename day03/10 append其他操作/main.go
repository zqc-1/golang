package main

import "fmt"

func main() {
	/*var s = []int{1, 2, 3}
	fmt.Printf("%p\n", &s)
	s = append(s, 4)
	fmt.Printf("%p\n", &s)*/

	//向开头插入值或者切片
	var a = []int{1, 2, 3}
	fmt.Println(append([]int{0}, a...))
	fmt.Println(append([]int{-1, 23, 444}, a...))

	//中间插入值
	var b = []int{1, 2, 3, 4}
	var index = 2
	fmt.Println(append(b[:index], append([]int{100}, b[index:]...)...))

	//删除
	var c = []int{1, 2, 3, 4, 5}
	c = append(c[:2], c[2+1:]...)
	fmt.Println(c)
}
