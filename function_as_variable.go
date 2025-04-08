package main

import "fmt"

func main() {

	fmt.Println("function as variable:")
	// function as variable
	var f func(int) int
	f = func(y int) int { return y + y }

	my_number := f(10)

	fmt.Println("my_number =", my_number)
}
