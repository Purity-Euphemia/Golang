package main

import "fmt"

func largestNumber(num []int) int{
	largestNumber := num[0]
	for count := 1; count < len(num); count++ {
		if num[count] > largestNumber {
			largestNumber = num[count]
		}
	}
	return largestNumber
}

func main() {
	numbers := [] int{10, 20 , 30, 40, 50}
	fmt.Println(largestNumber(numbers))
}