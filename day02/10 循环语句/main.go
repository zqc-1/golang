package main

import "fmt"

func main() {

	/*var count = 10
	for count > 0 {
		fmt.Println(count)
		count--
	}*/

	/*for i := 0; i < 10; i++ {
		fmt.Println(i)
	}*/

	var sum = 0
	for i := 0; i < 101; i++ {
		sum += i
	}
	fmt.Println(sum)
}
