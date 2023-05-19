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

// Student类型的方法接收器
func (s Student) read(book string) {
	fmt.Printf("%s学生正在读%s\n", s.name, book)
}

// Student类型的方法接收器
func (s Student) learn() {
	fmt.Printf("%s学生正在学习", s.name)
}

func main() {
	s1 := NewStudents(1001, "zqc", 32, []string{"chinese", "math", "english"})
	fmt.Println(s1.name)

	s1.read("射雕英雄传")
}
