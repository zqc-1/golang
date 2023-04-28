package main

import (
	"fmt"
)

func main() {
	//var names = []string{"zhu", "qian", "cheng"}
	//var ages = []int{18, 19, 20}
	//fmt.Println(names)
	//fmt.Println(ages)

	//var stus map[string]string
	//fmt.Println(stus, reflect.TypeOf(stus))

	//声明并初始化

	var stus = map[string]string{"name": "zhu", "age": "32"}
	//支持key查询
	fmt.Println(stus)
	fmt.Println(stus["name"])
	fmt.Println(stus["age"])
	fmt.Println(len(stus))
	//写入key-value
	stus["gender"] = "male"
	fmt.Println(stus) //map[age:32 gender:male name:zhu]
	stus["heigh"] = "180cm"
	//修改key-value
	stus["heigh"] = "190cm"
	fmt.Println(stus) //map[age:32 gender:male heigh:190cm name:zhu]
	//删除map
	delete(stus, "gender")
	fmt.Println(stus) //map[age:32 heigh:190cm name:zhu]

	//基于make函数声明初始化
	var stus1 = make(map[string]interface{})
	stus1["name"] = "zqc"
	stus1["age"] = 18
	stus1["gender"] = "male"
	fmt.Println(stus1)

	//遍历map
	for k, v := range stus1 {
		fmt.Println(k, v)
	}

	noSortMap := map[int]int{
		1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6,
	}
	for k, v := range noSortMap {
		fmt.Println(k, v)
	}
}
