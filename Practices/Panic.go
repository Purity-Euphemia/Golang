package main

import "fmt"

func safeDivide(a, b int) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered:", r)
        }
    }()

    if b == 0 {
        panic("cannot divide by zero")
    }

    fmt.Println("Result:", a/b)
}

func main() {
    safeDivide(10, 0)
}
