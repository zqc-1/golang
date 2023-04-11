package main

import "fmt"

func main() {
	//多分支语句 if - else if - else if ... else

	/**
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
	*/

	var week int
	fmt.Println("请输入星期数字")
	fmt.Scan(&week)

	/*if week == 1 {
		fmt.Println("星期一")
	} else if week == 2 {
		fmt.Println("星期二")
	} else if week == 3 {
		fmt.Println("星期三")
	} else if week == 4 {
		fmt.Println("星期四")
	} else if week == 5 {
		fmt.Println("星期五")
	} else if week == 6 {
		fmt.Println("星期六")
	} else if week == 7 {
		fmt.Println("星期日")
	} else {
		fmt.Println("输入错误")
	}*/

	//switch语句
	if week > 0 && week <= 7 {
		switch week {
		case 1:
			fmt.Println("星期一")
		case 2:
			fmt.Println("星期二")
		case 3:
			fmt.Println("星期三")
		case 4:
			fmt.Println("星期四")
		case 5:
			fmt.Println("星期五")
		case 6:
			fmt.Println("星期六")
		case 7:
			fmt.Println("星期日")
		default:
			fmt.Println("输入错误")
		}
	} else {
		fmt.Println("非法输入")
	}
}
