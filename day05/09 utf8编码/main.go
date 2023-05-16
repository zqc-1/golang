package main

import (
	"fmt"
)

func main() {
	var a rune
	a = '苑'
	fmt.Printf("unicode的十进制：%d\n", a)
	fmt.Printf("unicode的十六进制：%x\n", a)
	fmt.Printf("unicode的二进制：%b\n", a)

	var s = "苑abc"
	fmt.Println(s, len(s))

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	fmt.Println("=======")

	for i, v := range s {
		fmt.Println(i, v)
	}

}
