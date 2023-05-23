package main

import "fmt"

type Person struct {
	name string
}

func (per Person) eat() {
	fmt.Printf("%v,eat...\n", per.name)
}

func (per Person) sleep() {
	fmt.Printf("%v,Sleep...\n", per.name)
}

type Customer struct {
	name string
}

func (customer Customer) login(name string, pwd string) bool {
	fmt.Printf("customer.nane:%v\n", customer.name)
	if name == "tom" && pwd == "123" {
		return true
	} else {
		return false
	}
}

func main() {

	cus := Customer{
		name: "tom",
	}
	b := cus.login("tom", "123")
	fmt.Printf("%v", b)
	/*per := Person{
		name: "tom",
	}
	per.eat()
	per.sleep()*/
}
