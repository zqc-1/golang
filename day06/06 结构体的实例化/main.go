package main

import "fmt"

type Student struct {
	sid    int
	name   string
	age    int
	course []string
}

type Person struct {
	name   string
	age    int
	gender string
	addr   string
}

func main() {
	//声明方式一
	var stu Student
	stu.sid = 1001
	stu.name = "zqc"
	stu.age = 16
	stu.course = []string{"chinese", "math", "english"}
	fmt.Println(stu)
	fmt.Println(stu.name)
	stu.course = append(stu.course, "go")
	fmt.Println(stu)

	/*fmt.Printf("%p\n", &stu)
	fmt.Printf("%p\n", &(stu.sid))
	fmt.Printf("%p\n", &(stu.name))
	fmt.Printf("%p\n", &(stu.age))
	fmt.Printf("%p\n", &(stu.course))*/

	//声明方式二
	/*
			1、结构体可以使用“键值对”（Key value pair）初始化字段，每个“键”（Key）对应结构体中的一个字段，
		   键的“值”（Value）对应字段需要初始化的值。键值对的填充是可选的，不需要初始化的字段可以不填入初始化列表中，走默认值。

			2、多值初始化方式必须初始化结构体的所有字段且每一个初始值的填充顺序必须与字段在结构体中的声明顺序一致。

			3、键值对与值列表的初始化形式不能混用。
	*/
	stu2 := Student{sid: 1002, name: "zs", age: 18, course: []string{"chinese", "math", "english"}}
	fmt.Println(stu2)

	stu3 := Student{1003, "ls", 22, []string{"chinese", "math", "english"}}
	fmt.Println(stu3)
}
