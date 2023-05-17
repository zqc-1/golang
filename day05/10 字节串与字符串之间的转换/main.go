package main

import "fmt"

func main() {
	var msg = "abc苑"

	//字符串转换成字节串
	fmt.Println([]byte(msg))
	fmt.Println([]rune(msg))
	fmt.Println(len([]rune(msg)))

	//字节串转换成字符串
	info1 := []byte(msg)
	info2 := []byte{97, 98, 99, 232, 139, 145}
	fmt.Println(info1)
	fmt.Println(info2)
	fmt.Println(string(info1))
	fmt.Println(string(info2))

}
