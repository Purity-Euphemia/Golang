package main

import "fmt"

func main() {
	arr := [4]int{2, 4, 6, 8}
	sum := 0
	for _, v := range arr {
		sum += v
	}
	fmt.Println("Sum:", sum)
}
