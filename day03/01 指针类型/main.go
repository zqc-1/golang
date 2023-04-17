package main

import "fmt"

func main() {
	//&变量：获取变量地址
	/*var x = 10
	fmt.Printf("赋值之前的地址：%p\n", &x)
	x = 100
	fmt.Printf("赋值之后的地址：%p\n", &x)*/

	/*var x = 10
	fmt.Println(&x)

	var p *int //p是一个整形指针类型
	p = &x
	fmt.Println(p)

	//	取值操作 *指针变量
	fmt.Println(*p, reflect.TypeOf(*p))
	*p = 100
	fmt.Println(x)*/

	//var a = 1
	//var b = &a
	//var c = b
	//*b = 100
	//fmt.Println(a)
	//fmt.Println(*b)
	//fmt.Println(c)

	//指针应用1
	/*var x = 10
	var y = x
	var z = &x
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(*z)
	*z = 30
	fmt.Println(x)*/

	//指针应用2
	/*var x = 10
	var y = &x
	var z = *y
	x = 20
	fmt.Println(x) //20
	fmt.Println(*y) //20
	fmt.Println(z)// 10
	*y = 30
	fmt.Println(z)//10*/

	//var a = 100
	//var b = &a //*int
	//var c = &b //**int
	//fmt.Println(reflect.TypeOf(c))
	//**c = 200
	//fmt.Println(a)
	////fmt.Println(b) //a的地址
	////fmt.Println(c) //b的地址

	p1 := 1
	p2 := &p1
	*p2++
	fmt.Println(p1)
	fmt.Println(*p2)
}
