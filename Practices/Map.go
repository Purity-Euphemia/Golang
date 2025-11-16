package main

import "fmt"

func main() {
    ages := map[string]int{
        "Alice": 25,
        "Bob":   30,
    }

    ages["Charlie"] = 22

    fmt.Println("Alice is", ages["Alice"])
    fmt.Println("All data:", ages)
}
