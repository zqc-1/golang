package main

import (
	"fmt"
	"reflect"
)

/*func getCounter() func() {
	var i = 0

	counter := func() {
		i++
		fmt.Println(i)
		fmt.Println(&i)
	}
	return counter
}*/

func getCounter(i int) func() {

	counter := func() {
		i++
		fmt.Println(i)
		fmt.Println(&i)
	}
	return counter
}

func main() {
	counter1 := getCounter(10)
	fmt.Println(reflect.TypeOf(counter1))
	counter1()
	counter1()
	counter1()
	counter2 := getCounter(5)
	counter2()
	counter2()
}
