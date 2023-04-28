package main

import (
	"fmt"
	"strconv"
)

func main() {
	//map嵌套slice
	/*var data = make(map[string][]string)
	data["beijing"] = []string{"朝阳", "海淀", "昌平"}
	data["henan"] = []string{"郑州", "开封", "洛阳"}
	data["shanghai"] = []string{"徐汇", "浦东", "静安"}
	fmt.Println(data)*/
	//查找河南的第二个城市
	//fmt.Println(data["henan"][1])
	//遍历每一个省份以及对应的成市
	/*for proStr, citySlices := range data {
		//fmt.Println(proStr,citySlices)
		fmt.Println(proStr, len(citySlices))
		for i, city := range citySlices {
			fmt.Printf("%d.%s ", i, city)
		}
		fmt.Println()
	}*/

	//删除北京的k-v
	/*delete(data, "beijing")
	fmt.Println(data)*/

	//map嵌套map
	/*stu01 := map[string]string{"name": "zhu", "age": "18"}
	stu02 := map[string]string{"name": "qian", "age": "19"}
	stu03 := map[string]string{"name": "cheng", "age": "20"}
	var stus = make(map[int]map[string]string)
	stus[1001] = stu01
	stus[1002] = stu02
	stus[1003] = stu03
	fmt.Println(stus)
	fmt.Println(len(stus))
	//打印学生1002的年龄
	fmt.Println(stus[1002]["age"])
	//循环打印每一个学生的学号姓名年龄
	for sid, stu := range stus {
		fmt.Println(sid, stu["name"], stu["age"])
	}*/

	//切片嵌套map
	//方式一：
	stu01 := map[string]string{"name": "zhu", "age": "18"}
	stu02 := map[string]string{"name": "qian", "age": "19"}
	stu03 := map[string]string{"name": "cheng", "age": "20"}
	//var stus = make([]map[string]string, 3)
	var stus = []map[string]string{stu01, stu02, stu03}
	//方式二：
	//var stus = []map[string]string{{"name": "zhu", "age": "18"}, {"name": "qian", "age": "19"}, {"name": "cheng", "age": "20"}}

	//打印第二个学生的姓名
	fmt.Println(stus[1]["name"])
	//循环打印每一个学生的学号姓名年龄
	for _, stu := range stus {
		//fmt.Println(stu["name"], stu["age"])
		fmt.Printf("姓名：%-6s 年龄：%-6s\n", stu["name"], stu["age"])
	}

	//添加一个学生对象

	/*var name, age string
	fmt.Println("请输入学生的姓名及年龄“")
	fmt.Scan(&name, &age)
	newMap := map[string]string{"name": name, "age": age}
	stus = append(stus, newMap)
	fmt.Println(stus)*/

	//删除一个学生qian的map
	// 查询qian的索引位置

	/*var delIndex = 0
	for index, stuMap := range stus {
		if stuMap["name"] == "cheng" {
			delIndex = index
		}
	}
	stus = append(stus[:delIndex], stus[delIndex+1:]...)
	fmt.Println(stus)*/

	//将姓名为qian的年龄自动加1
	for _, stuMap := range stus {
		if stuMap["name"] == "zhu" {
			//类型转换
			age, _ := strconv.Atoi(stuMap["age"])
			stuMap["age"] = strconv.Itoa(age + 1)
		}
	}
	fmt.Println(stus)
}
