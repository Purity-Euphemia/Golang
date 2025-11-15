package main

import "fmt"

func main() {
    ch := make(chan int, 1)

    select {
    case val := <-ch:
        fmt.Println(val)
    default:
        fmt.Println("no message")
    }
}
