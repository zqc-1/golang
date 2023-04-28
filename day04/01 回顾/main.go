package main

import "fmt"

func main() {

	/*var x = 10
	//指针变量
	var p *int
	p = &x
	fmt.Println(p)

	//取值
	fmt.Println(*p)
	*p = 100*/

	//=====================

	/*var p *int
	p = new(int)
	fmt.Println(*p)*/

	/*var arr [3]int
	fmt.Println(arr)

	fmt.Printf("%p\n", &arr)
	fmt.Println(&arr[0])
	fmt.Println(&arr[1])
	fmt.Println(&arr[2])*/

	//切片
	var arr = [5]int{1, 2, 3, 4, 5}
	s1 := arr[0:3]
	s2 := arr[2:5]
	s3 := s2[0:2]

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

}
