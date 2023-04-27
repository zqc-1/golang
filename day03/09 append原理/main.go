package main

import "fmt"

func main() {
	/*a := []int{11, 22, 33}
	fmt.Println(len(a), cap(a))

	c := append(a, 44)
	a[0] = 100
	fmt.Println(a)
	fmt.Println(c)*/

	//make
	/*a := make([]int, 3, 10)
	fmt.Println(a)
	b := append(a, 11, 22)
	fmt.Println(a)
	fmt.Println(b)
	a[0] = 100
	fmt.Println(a)
	fmt.Println(b)*/

	arr := [4]int{10, 20, 30, 40}
	s1 := arr[0:2]
	s2 := s1
	s3 := append(append(append(s1, 1), 2), 3)
	s1[0] = 100
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(arr)
}
