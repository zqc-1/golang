package main

import "fmt"

/*
	type Student struct {
		string //string string
		int    //int int
	}
*/
type Addr struct {
	province string
	city     string
	town     string
}
type Student struct {
	name string
	age  int
	//addr Addr
	Addr
}

func main() {
	s := Student{"zqc", 22, Addr{"河南", "郑州", "二七"}}
	fmt.Println(s.Addr)
	fmt.Println(s.Addr.town)
	fmt.Println(s.town)
}
