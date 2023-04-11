package main

import (
	"fmt"
	"os"
)

func main() {
	var month, day int
	fmt.Print("请输入生日的月和日：")
	fmt.Scan(&month, &day)

	var xingZuo string

	if day < 1 || day > 31 {
		fmt.Println("输入的日有问题")
		os.Exit(0)
	}

	switch month {
	case 1:
		// 日判断
		if day >= 1 && day <= 19 {
			xingZuo = "摩羯座"
		} else {
			xingZuo = "水瓶座"
		}

	case 2:
		// 日判断
		if day >= 1 && day <= 18 {
			xingZuo = "水瓶座"
		} else {
			xingZuo = "双鱼座"
		}
	case 3:
		// 日判断
		if day >= 1 && day <= 20 {
			xingZuo = "双鱼座"
		} else {
			xingZuo = "白羊座"
		}
	case 4:
		// 日判断
		if day >= 1 && day <= 19 {
			xingZuo = "白羊座"
		} else {
			xingZuo = "金牛座"
		}
	case 5:
		// 日判断
		if day >= 1 && day <= 20 {
			xingZuo = "金牛座"
		} else {
			xingZuo = "双子座"
		}
	case 6:
		// 日判断
		if day >= 1 && day <= 21 {
			xingZuo = "双子座"
		} else {
			xingZuo = "巨蟹座"
		}
	case 7:
		// 日判断
		if day >= 1 && day <= 22 {
			xingZuo = "巨蟹座"
		} else {
			xingZuo = "狮子座"
		}
	case 8:
		// 日判断
		if day >= 1 && day <= 22 {
			xingZuo = "狮子座"
		} else {
			xingZuo = "处女座"
		}
	case 9:
		// 日判断
		if day >= 1 && day <= 22 {
			xingZuo = "处女座"
		} else {
			xingZuo = "天秤座"
		}
	case 10:
		// 日判断
		if day >= 1 && day <= 23 {
			xingZuo = "天秤座"
		} else {
			xingZuo = "天蝎座"
		}
	case 11:
		// 日判断
		if day >= 1 && day <= 22 {
			xingZuo = "天蝎座"
		} else {
			xingZuo = "射手座"
		}
	case 12:
		// 日判断
		if day >= 1 && day <= 21 {
			xingZuo = "射手座"
		} else {
			xingZuo = "摩羯座"
		}
	default:
		fmt.Println("输入的月份有问题")
	}

	fmt.Println("您的星座是：", xingZuo)
}
