package main

import "fmt"

func main() {
	//方式一：通过数组切片操作获得切片对象
	//var arr = [3]string{"zhu", "qian", "cheng"}
	//s1 := arr[0:2]
	//fmt.Println(s1, reflect.TypeOf(s1))
	//s2 := arr[1:]
	//fmt.Println(s2, reflect.TypeOf(s2))
	//
	//s2[0] = "zqc"
	//fmt.Println(s1)
	//fmt.Println(arr)

	var s = []int{1, 2, 3, 4, 5}
	s1 := s[1:4]
	fmt.Println(len(s1), cap(s1))
}
