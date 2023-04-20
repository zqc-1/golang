package main

import (
	"fmt"
	"reflect"
)

func main() {

	var names = [3]string{"zhu", "qian", "cheng"}
	// 索引操作
	fmt.Println(names[2])
	names[2] = "Cheng"
	fmt.Println(names[2])

	//切片操作
	var arr = [5]int{11, 12, 13, 14, 15}
	s1 := arr[1:3]
	fmt.Println(s1, reflect.TypeOf(s1))
	s2 := arr[1:]
	fmt.Println(s2, reflect.TypeOf(s2))
	s3 := arr[:3]
	fmt.Println(s3, reflect.TypeOf(s3))

	//遍历数组
	for i := 0; i < len(arr); i++ {
		fmt.Println(i, arr[i])
	}

	fmt.Println("============================================")

	//range循环
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
