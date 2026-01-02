package main

import "fmt"

func sumArray(arr [5]int) int {
    total := 0
    for _, value := range arr {
        total += value
    }
    return total
}

func main() {
    numbers := [5]int{1, 2, 3, 4, 5}
    result := sumArray(numbers)
    fmt.Println("Sum:", result)
}
