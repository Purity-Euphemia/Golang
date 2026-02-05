package main

import "fmt"

func main() {
	arr := [6]int{1, 2, 3, 4, 5, 6}
	count := 0
	for _, v := range arr {
		if v%2 == 0 {
			count++
		}
	}
	fmt.Println("Even count:", count)
}
