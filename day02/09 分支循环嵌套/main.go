package main

import "fmt"

func main() {
	var username, password string
	fmt.Println("请输入用户名")
	fmt.Scan(&username)
	fmt.Println("请输密码")
	fmt.Scan(&password)

	if username == "root" && password == "root" {
		var score int
		fmt.Print("请输入考试成绩：")
		fmt.Scan(&score)

		if score < 0 || score > 100 {
			fmt.Println("非法数字")
		} else if score >= 90 {
			fmt.Println("A")
		} else if score > 80 {
			fmt.Println("B")
		} else if score > 60 {
			fmt.Println("C")
		} else {
			fmt.Println("D")
		}
	} else {
		fmt.Println("用户名或密码错误")
	}
}
