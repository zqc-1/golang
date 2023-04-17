package main

import "fmt"

func main() {

	//基本数据类型属于值类型
	//值类型：（整形，浮点型、布尔类型、数组、结构体）当声明未赋值之前存在一个默认值
	var x int
	fmt.Println(x)

	var name string
	fmt.Println(name)

	//指针类型属于引用类型	包括切片 map channel都属于引用类型

	var p *int
	p = new(int)
	fmt.Println(*p)
	*p = 10
	fmt.Println(*p)
}
