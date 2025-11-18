package main

import (
    "fmt"
    "time"
)

func worker(id int, jobs <-chan int) {
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
        time.Sleep(300 * time.Millisecond)
    }
}

func main() {
    jobs := make(chan int)

    go worker(1, jobs)
    go worker(2, jobs)

    for i := 1; i <= 5; i++ {
        jobs <- i
    }

    close(jobs)
}
