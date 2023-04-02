package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	//整数之间的转换
	var x int8 = 10
	var y int16 = 20
	fmt.Println(x + int8(y))
	//var s = "100"
	//fmt.Println(s, reflect.TypeOf(s))

	//字符串与整形之间的转换
	var ageStr = "22"
	var age, err = strconv.Atoi(ageStr)
	fmt.Println("err", err) //<nil> 空
	fmt.Println(age)

	price := 100
	var priceStr = strconv.Itoa(price)
	fmt.Println(priceStr, reflect.TypeOf(priceStr))

	//strconv parse系函数
	//将数字字符串转换为整型
	ret, _ := strconv.ParseInt("22", 10, 8)
	fmt.Println(ret, reflect.TypeOf(ret))
	//将数字字符串转换为浮点型
	ret2, _ := strconv.ParseFloat("3.1415926", 64)
	fmt.Println(ret2, reflect.TypeOf(ret2))
	//将数字字符串转换为布尔值
	// 返回字符串表示的bool值。它接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE；否则返回错误。
	ret3, _ := strconv.ParseBool("1")
	fmt.Println(ret3, reflect.TypeOf(ret3))
}
