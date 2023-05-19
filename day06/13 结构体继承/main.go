package main

import "fmt"

// Animal 动物
type Animal struct {
	name string
}

func (a *Animal) eat() {
	fmt.Printf("%s is eating！\n", a.name)
}
func (a *Animal) sleep() {
	fmt.Printf("%s is sleeping！\n", a.name)
}

// Dog 类型
type Dog struct {
	Kind    string
	*Animal //通过嵌套匿名结构体实现继承
}

func (d *Dog) bark() {
	fmt.Printf("%s is barking ~\n", d.name)
}

// Cat 类型
type Cat struct {
	*Animal
}

func (c *Cat) climbTree() {
	fmt.Printf("%s is climb tree ~\n", c.name)
}
func main() {

	d1 := &Dog{
		Kind: "金毛",
		Animal: &Animal{ //注意嵌套的是结构体指针
			name: "旺财",
		},
	}
	fmt.Println(d1.Animal.name)
	d1.eat()
	d1.bark()
	d1.sleep()

	c1 := &Cat{
		Animal: &Animal{
			name: "喵喵",
		},
	}
	c1.sleep()
	c1.climbTree()
}
