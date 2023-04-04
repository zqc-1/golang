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

	//var (
	//	name      string
	//	age       int
	//	isMarried bool
	//)
	//fmt.Scanf("1:%s 2:%d 3:%t", &name, &age, &isMarried)
	//fmt.Printf("扫描结果 姓名:%s 年龄:%d 婚否:%t", name, age, isMarried)

	//var name string
	//var age int
	//fmt.Print("请输入姓名和年龄：")
	//fmt.Scanln(&name, &age)
	//fmt.Println(name)
	//fmt.Println(age)

	var a, b int
	fmt.Printf("请输入两个数字：")
	fmt.Scanf("%d+%d", &a, &b)
	fmt.Println(a + b)

}
