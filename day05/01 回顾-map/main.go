package main

import "fmt"

func main() {
	/*var info = make(map[string]string)
	info["name"] = "zqc"
	info["age"] = "22"
	fmt.Println(info)
	info2 := info
	info2["gender"] = "male"
	fmt.Println(info2)
	info["height"] = "180cm"
	fmt.Println(info2) // map[age:22 gender:male height:180cm name:zqc]*/

	/*var info = make(map[string]string)
	info["name"] = "zqc"
	info["age"] = "22"
	info["gender"] = "male"
	fmt.Println(info["age"])
	value, exist := info["gender"]
	fmt.Println(value, exist)
	if exist {
		fmt.Println("性别存在")
	} else {
		fmt.Println("性别不存在")
	}
	delete(info, "gender")
	for k, v := range info {
		fmt.Println(k, v)
	}*/

	/*stu01 := map[string]string{"name": "zhu", "age": "18"}
	stu02 := map[string]string{"name": "qian", "age": "19"}
	stu03 := map[string]string{"name": "cheng", "age": "20"}
	stus := []map[string]string{stu01, stu02, stu03}
	fmt.Println(stus[1]["age"])
	//var stu = stus[1]
	//fmt.Println(stu["age"])*/

	//map嵌套map
	/*var data = make(map[string]map[string]string)
	data["zhu"] = map[string]string{"age": "18", "gender": "male"}
	data["qian"] = map[string]string{"age": "18", "gender": "female"}
	fmt.Println(data["qian"]["gender"])*/

	var data = make(map[string][]string)
	data["山东省"] = []string{"威海", "日照"}
	fmt.Println(data)
}
