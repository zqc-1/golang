package main

import "fmt"

func printRange(start, end int) {
	for i := start; i <= end; i++ {
		fmt.Println(i)
	}
}

func add(x, y int) int {
	z := x + y
	return z
}
func main() {
	fmt.Println("Hello world!")
	printRange(1, 10)
	fmt.Println("Hello world!")

}
