package main

import "fmt"

func main() {
	numbers := [5]int{10, 20, 30, 40, 50}

	fmt.Println("Array:", numbers)

	fmt.Println("First element:", numbers[0])
	fmt.Println("Last element:", numbers[4])

	for i, value := range numbers {
		fmt.Println("Index", i, "Value", value)
	}
}
