package main

import "fmt"

func minMax(a, b int) (int, int) {
    if a < b {
        return a, b
    }
    return b, a
}

func main() {
    low, high := minMax(3, 7)
    fmt.Println("Min:", low, "Max:", high)
}
