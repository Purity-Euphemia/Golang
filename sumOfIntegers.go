package main

import "fmt"

func sumOfVariable(a, b int) int {
	c := a + b
	return c
}

func main() {
	fmt.Println(sumOfVariable(10, 5))
}