package main

import "fmt"

func main() {

	//案例一
	/*fmt.Println("test01")
	defer fmt.Println("test02")
	fmt.Println("test03")*/

	//案例二
	/*f, err := os.Open("满江红")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close() //关闭文件*/

	//案例三
	/*fmt.Println("test01")
	defer fmt.Println("test02")
	fmt.Println("test03")
	defer fmt.Println("test04")
	fmt.Println("test05")*/

	//案例四
	/*foo := func() {
		fmt.Println("I am function foo1")
	}
	defer foo()
	foo = func() {
		fmt.Println("I am function foo2")
	}*/

	//案例五
	/*x := 10
	defer func(a int) {
		fmt.Println(a)
	}(x)
	x++*/

	//案例六
	x := 10
	defer func() {
		fmt.Println(x)
	}()
	x++

}
