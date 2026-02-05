package main

import "fmt"

func main() {
	arr := [5]int{3, 9, 1, 7, 4}
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	fmt.Println("Max:", max)
}
