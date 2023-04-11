package main

import (
	"fmt"
)

func main() {

	//单分支 if
	//var name = "zqc"
	//if name == "root" {
	//	fmt.Println("匹配成功")
	//}
	//fmt.Println("end")

	//双分支语句 if else
	/**
	var age int
	fmt.Println("请输入年龄：")
	fmt.Scan(&age)
	if age < 18 {
		fmt.Println("未满18岁，审核不通过，不能玩改游戏！")
	} else {
		fmt.Println("审核通过！")
	}
	*/
	var username, password string
	fmt.Println("请输入用户名")
	fmt.Scan(&username)
	fmt.Println("请输密码")
	fmt.Scan(&password)

	if username == "root" && password == "root" {
		fmt.Println("欢迎进入")
	} else {
		fmt.Println("用户名或密码错误")
	}
}
