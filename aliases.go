package main

import "fmt"

func main() {
	fmt.Println("alias:")
	type Bytesize int64
	type datasize float64
	var b Bytesize = 1024
	var c datasize = 1024
	fmt.Println("b =", b)
	fmt.Println("c =", c)
}
