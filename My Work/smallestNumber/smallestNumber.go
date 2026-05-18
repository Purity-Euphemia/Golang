package main

import "fmt"

func main() {
	nums := [] int{10, 20, 30, 40, 50}
	smallestNumber := nums[0]
	for count := 1; count < len(nums); count++ {
		if nums[count] < smallestNumber{
			smallestNumber = nums[count]
		}
	}
	fmt.Println(smallestNumber)
}