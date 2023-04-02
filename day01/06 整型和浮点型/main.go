package main

import (
	"fmt"
	"reflect"
)

func main() {

	//var x int8 = 100
	//fmt.Println(x)

	//var f1 float32 = 2.123456789123456
	//var f2 float64 = 2.123456789123456
	//fmt.Println(f1, reflect.TypeOf(f1))
	//fmt.Println(f2, reflect.TypeOf(f2))

	//var f3 = 2e10
	//fmt.Println(f3, reflect.TypeOf(f3))

	//var b bool
	//b = true
	//b = false

	//c := 2 > 1
	//fmt.Println(c, reflect.TypeOf(c))
	//
	//var x uint8 = 25
	//fmt.Println(x)

	var f1 float32 = 1.123456789
	var f2 float64 = 1.12345678912345678
	fmt.Println(f1, reflect.TypeOf(f1))
	fmt.Println(f2, reflect.TypeOf(f2))
}
