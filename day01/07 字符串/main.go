package main

import "fmt"

func main() {

	var s string
	s = "hello world"

	//取值
	fmt.Println(s)
	fmt.Println(string(s[1]))
	fmt.Println(string(s[10]))

	//切片
	fmt.Println(string(s[0:5]))
	fmt.Println(string(s[:5]))
	fmt.Println(string(s[6:11]))
	fmt.Println(string(s[:]))

	//拼接
	var s1 = "123"
	var s2 = " 456"
	fmt.Println(s1 + s2)

	//转义符 \
	fmt.Println("123\n456")

	var s3 = "D:\\Learn\\Go\\hello.exe"
	fmt.Println(s3)

	var s4 = "his name is \"jack\""
	fmt.Println(s4)

	//多行打印
	fmt.Println("1. 赚金币")
	fmt.Println("2. 买武器")
	fmt.Println("3. 买血")

	info := `	
1. 赚金币	
2. 买武器
3. 买血
请选择：
	`
	fmt.Println(info)
}
