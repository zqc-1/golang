package main

import "fmt"

func main() {

	//数组必须限制长度
	//声明
	var arr [3]string
	fmt.Println(arr)

	//赋值
	//arr[0] = "zhu"
	//arr[1] = "qian"
	//arr[2] = "cheng"
	//fmt.Println(arr)

	//数组声明并赋值
	var names = [3]string{"zhu", "qian", "cheng"}
	fmt.Println(names)
	//for i := 0; i < len(names); i++ {
	//	fmt.Println(names[i])
	//}

	//省略长度赋值
	var names2 = [...]string{"11", "22", "33"}
	fmt.Println(names2)

	//索引复制
	var names3 = [...]string{0: "zhu", 2: "cheng"}
	fmt.Println(names3)

	//go len函数：计算容器数据长度
	fmt.Println(len("hello"))
	fmt.Println(len(names))
}
