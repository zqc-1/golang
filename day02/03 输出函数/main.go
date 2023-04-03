package main

import "fmt"

func main() {
	//name := "zqc"
	//age := 24
	//
	//fmt.Print(name, age)
	//
	//fmt.Println("hello world")
	//fmt.Println(name)
	//fmt.Println(age)
	//fmt.Println(name, age)
	//fmt.Println("姓名:", name, "年龄:", age)

	//name := "yuan"
	//age := 24
	//isMarried := false
	//salary := 3000.549
	//fmt.Printf("姓名:%s 年龄:%d 婚否：%t 薪资:%.2f\n", name, age, isMarried, salary)
	//fmt.Printf("姓名:%v 年龄:%v 婚否：%v 薪资:%v\n", name, age, isMarried, salary)
	//fmt.Printf("姓名:%#v 年龄:%#v 婚否：%#v 薪资:%#v\n", name, age, isMarried, salary)

	// 整形和浮点型
	fmt.Printf("%b\n", 12)       // 二进制表示:1100
	fmt.Printf("%d\n", 12)       // 十进制表示:12
	fmt.Printf("%o\n", 12)       // 八进制表示:14
	fmt.Printf("%x\n", 12)       // 十六进制表示:c
	fmt.Printf("%X\n", 12)       // 十六进制表示：c
	fmt.Printf("%f\n", 3.1415)   // 有小数点而无指数：3.141500
	fmt.Printf("%.3f\n", 3.1415) // 3.142
	fmt.Printf("%e\n", 1000.25)  // 科学计数法:  1.000250e+03，默认精度为6

	// 设置宽度
	fmt.Printf("学号：%s 姓名：%-20s 平均成绩：%-4d\n", "1001", "alvin", 100)
	fmt.Printf("学号：%s 姓名：%-20s 平均成绩：%-4d\n", "1002", "zuibangdeyuanlaoshi", 98)
	fmt.Printf("学号：%s 姓名：%-20s 平均成绩：%-4d\n", "1003", "x", 78)

}
