package main

import "fmt"

func doThing() (int, error) {
    return 42, nil
}

func main() {
    val, err := doThing()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("Value:", val)
}
