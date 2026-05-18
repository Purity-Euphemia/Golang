package main

import "fmt"

func averageOfVariable(array []int) int {
	sum := 0
	for count := 0; count < len(array); count++ {
		sum = sum + array[count]
	}
	return sum / len(array)
}

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println(averageOfVariable(numbers))
}
