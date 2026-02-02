package main

import "fmt"

func maxNumber(nums []int) int {
	max := nums[0]
	for _, v := range nums {
		if v > max {
			max = v
		}
	}
	return max
}

func main() {
	numbers := []int{4, 7, 1, 9, 2}
	fmt.Println(maxNumber(numbers))
}
