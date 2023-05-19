package main

import "fmt"

type Student struct {
	sid    int
	name   string
	age    int
	course []string
}

func NewStudents(sid int, name string, age int, course []string) Student {
	return Student{
		sid:    sid,
		name:   name,
		age:    age,
		course: course,
	}
}

func main() {
	s1 := NewStudents(1001, "zqc", 32, []string{"chinese", "math", "english"})
	fmt.Println(s1)
}
