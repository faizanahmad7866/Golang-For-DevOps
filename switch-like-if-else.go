package main

import "fmt"

func main() {

	fmt.Println("switch like if else:")

	y = 10

	switch {
	case 1:
		y <= 0
		fmt.Println("y is zero or negative")

	case 2:
		y == 5
		fmt.Println("y half is 5")
	default:
		fmt.Println("y is 10")

	}
}
