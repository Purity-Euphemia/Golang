package main

import "fmt"

func Sum[T int | float64](nums []T) T {
    var total T
    for _, v := range nums {
        total += v
    }
    return total
}

func main() {
    fmt.Println(Sum([]int{1, 2, 3}))
    fmt.Println(Sum([]float64{1.5, 2.5}))
}
