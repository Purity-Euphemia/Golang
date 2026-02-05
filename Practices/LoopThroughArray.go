package main

import "fmt"

func main() {
	arr := [5]int{5, 10, 15, 20, 25}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
