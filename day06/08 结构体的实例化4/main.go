package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	sid    int
	name   string
	age    int
	course []string
}

func initSid(p *Student) {
	p.sid = 0
}

func main() {
	var s1 = new(Student)

	/*fmt.Println(s1)
	fmt.Println(s1.name)
	initSid(s1)
	fmt.Println(s1)
	fmt.Println(s1.sid)*/

	fmt.Println(reflect.TypeOf(s1))

	var s2 = make([]int, 3)
	fmt.Println(reflect.TypeOf(s2))

	var s3 = make(map[string]string, 3)
	fmt.Println(reflect.TypeOf(s3))
}
