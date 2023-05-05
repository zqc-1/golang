package main

import "fmt"

func setAge(p *int) {
	fmt.Println(p)
	fmt.Println(*p)
	*p++

}
func main() {
	var age = 18
	//var p = &age
	setAge(&age)
	fmt.Println(age)
}
