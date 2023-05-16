package main

import (
	"fmt"
	"reflect"
)

func main() {

	var y byte
	y = 'A'
	fmt.Println(y, reflect.TypeOf(y))
	fmt.Printf("%c\n", y)
	fmt.Printf("%d\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)

	var x uint8
	x = 65
	fmt.Printf("%c\n", x)
	fmt.Printf("%d\n", x)
	fmt.Println(x)

	//rune类型
	var z rune
	z = '朱'
	fmt.Println(z, string(z))
}
