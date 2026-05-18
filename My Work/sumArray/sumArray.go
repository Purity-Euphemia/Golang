package main

import "fmt"

func addArray(nums []int) int {
	sum := 0
	for count := 0; count < len(nums); count++{
		sum = sum + nums[count]
	}
	return sum
}

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Println(addArray(numbers))
}