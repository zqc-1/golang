package main

import (
	"fmt"
	"strings"
)

func main() {

	var name = "Zhuqc"
	var newName = strings.ToUpper(name)
	fmt.Println(newName)
	fmt.Println(strings.ToUpper(name))
	fmt.Println(strings.ToLower(name))

	var s = "zhu qian cheng"
	fmt.Println(strings.HasPrefix(s, "zhu"))
	fmt.Println(strings.HasSuffix(s, "g"))
	fmt.Println(strings.Contains(s, "q"))
	fmt.Println(strings.Contains(s, " q1"))

	username := " zhu "
	fmt.Println(strings.Trim(username, " "))
	username = strings.Trim(username, " ")
	fmt.Println(username == "zhu")

	//index 索引
	var s2 = "zhu qian cheng"
	fmt.Println(strings.Index(s2, "q"))

	var cmd = "mysql ... -u root -p 123"
	uIndex := strings.Index(cmd, "-u")
	pIndex := strings.Index(cmd, "-p")
	fmt.Println(uIndex)
	fmt.Println(pIndex)
	fmt.Println(cmd[uIndex+3 : pIndex-1])
	fmt.Println(cmd[pIndex+3:])

	//分割 拼接
	var s3 = "zhu qian cheng"
	nameSlice := strings.Split(s3, " ")
	fmt.Println(nameSlice)
	fmt.Println(nameSlice[0])

	newStr := strings.Join(nameSlice, ",")
	fmt.Println(newStr)
}
