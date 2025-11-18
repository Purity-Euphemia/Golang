package main

import "fmt"

func main() {
    students := map[string][]int{
        "Alice": {90, 80, 100},
        "Bob":   {70, 85, 90},
    }

    for name, scores := range students {
        fmt.Println(name, "scores:", scores)
    }
}
