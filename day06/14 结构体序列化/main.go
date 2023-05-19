package main

import (
	"encoding/json"
	"fmt"
)

type Addr struct {
	Province string
	City     string
}

type Stu struct {
	Name string `json:"name"` // 结构体的标签
	Age  int    `json:"-"`    // 表示不参与序列化
	Addr Addr
}

func main() {
	var stuMap = map[string]interface{}{"name": "zzz", "age": 22, "addr": "beijing"}
	var stuStruct = Stu{Name: "zzz", Age: 18, Addr: Addr{Province: "Hebei", City: "langFang"}}

	jsonStuMap, _ := json.Marshal(stuMap)
	fmt.Println(string(jsonStuMap))

	jsonStuStruct, _ := json.Marshal(stuStruct)
	fmt.Println(string(jsonStuStruct))

	//反序列化
	var StuMap map[string]interface{}
	err := json.Unmarshal(jsonStuMap, &StuMap)
	if err != nil {
		return
	}
	fmt.Println(StuMap, stuMap["name"])

	var StuStruct Stu
	err = json.Unmarshal(jsonStuStruct, &StuStruct)
	if err != nil {
		return
	}
	fmt.Println(StuStruct, StuStruct.Name)

	/*var s1 = Stu{Name: "s1", Age: 20, Addr: Addr{Province: "Hebei", City: "langFang"}}
	var s2 = Stu{Name: "s2", Age: 21, Addr: Addr{Province: "Hebei", City: "langFang"}}
	var s3 = Stu{Name: "s3", Age: 22, Addr: Addr{Province: "Hebei", City: "langFang"}}

	var data = []Stu{s1, s2, s3}
	jsonData, _ := json.Marshal(data)
	fmt.Println(string(jsonData))

	var data2 []Stu
	json.Unmarshal(jsonData, &data2)
	fmt.Println(data2, data2[1].Name)*/
}
