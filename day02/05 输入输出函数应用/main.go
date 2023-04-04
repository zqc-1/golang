package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("请输入您的生日，按格式: 年-月-日：")
	var birthday string
	fmt.Scan(&birthday)
	//fmt.Println("birthday", birthday)
	birth := strings.Split(birthday, "-")
	fmt.Println("birth", birth)
	fmt.Printf("您的生日是%s年-%s月-%s日", birth[0], birth[1], birth[2])

	var name string
	fmt.Println("请输入与个英文名：")
	fmt.Scan(&name)
	upperName := strings.ToUpper(name)
	var b = strings.HasPrefix(upperName, "A")
	fmt.Println(b)
}
