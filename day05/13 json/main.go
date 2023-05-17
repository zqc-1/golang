package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
)

func main() {
	var s = []int{1, 2, 3}
	fmt.Println(s)
	var m = map[string]string{"name": "zqc", "age": "22"}
	//json序列化
	data, _ := json.Marshal(m)
	//fmt.Println(string(data))
	//fmt.Printf("%#v", string(data))
	ioutil.WriteFile("data.json", data, 0666)

	//反序列化
	res, _ := ioutil.ReadFile("data.json")
	fmt.Println(res, reflect.TypeOf(res))
	fmt.Println(string(res))
	var info = make(map[string]string)
	json.Unmarshal(res, &info)
	fmt.Println(info, reflect.TypeOf(info))
}
