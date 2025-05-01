/*if Statement :--
The if statemnent in go is ues dor conditional statement excuation of code blocks,it can optinally include an else block for alternative execution when the if condition is not meet*/

package main

import "fmt"

func main() {
	fmt.Println("if--else:")
	x := 10

	if x > 5 {
		fmt.Println("x is greater than 5")

	} else {
		fmt.Println("x is not greater than 5")
	}

	fmt.Println("if else if else:")

	if x > 5 {
		fmt.Println("x is greater than 5")
	} else if x > 3 {
		fmt.Println("x is greater than 3 but less than or equal to 5")
	} else {
		fmt.Println("x is not greater than 5")
	}

	fmt.Println("With Initalization Statement:")
	if y := 20; y > 10 {
		fmt.Println("y is greater than 10")
	}
	// Note : y is not accessible outside the if block

}
