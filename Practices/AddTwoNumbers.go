package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main(){
	sum := add(4, 3)
	fmt.Println("sum", sum)
}