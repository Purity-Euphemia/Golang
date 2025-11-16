package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan int)

    go func() {
        time.Sleep(2 * time.Second)
        ch <- 10
    }()

    select {
    case v := <-ch:
        fmt.Println("Received:", v)
    case <-time.After(1 * time.Second):
        fmt.Println("Timeout — no message received")
    }
}
