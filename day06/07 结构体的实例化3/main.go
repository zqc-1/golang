package main

import "fmt"

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

	/*var s1 = Student{name: "zqc", sid: 1001}
	fmt.Println(s1)

	var s2 = s1
	s2.age = 32
	//fmt.Println(s2)
	fmt.Println(s1.age)

	initSid(&s1)
	fmt.Println(s1.sid)*/

	var s1 = &Student{name: "zqc", sid: 1001}
	fmt.Println(s1)
	initSid(s1)
	fmt.Println(s1.sid)
}
