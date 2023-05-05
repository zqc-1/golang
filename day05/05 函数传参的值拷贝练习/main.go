package main

import "fmt"

func foo(s []int) {
	s[1] = 200
	s = append(s, 4)
	s[0] = 100
	fmt.Println(s, len(s), cap(s)) // [100 200 3 4] 4 6
}

func main() {

	var s = []int{1, 2, 3}
	foo(s)
	fmt.Println(s)
}
