package main

import "fmt"

func main() {
	//var (
	//	name      string
	//	age       int
	//	isMarried bool
	//)
	//fmt.Scan(&name, &age, &isMarried)
	//fmt.Printf("扫描结果 name:%s age:%d married:%t\t", name, age, isMarried)

	// 案例1
	//var (
	//	name      string
	//	age       int
	//	isMarried bool
	//)
	//fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &isMarried)
	//fmt.Printf("扫描结果 姓名:%s 年龄:%d 婚否:%t", name, age, isMarried)

	// 案例2
	var a, b int
	fmt.Scanf("%d+%d", &a, &b)
	fmt.Println(a + b)
}
