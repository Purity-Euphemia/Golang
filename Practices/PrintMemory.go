package piscine

import "fmt"

func PrintMemory(arr [10]byte) {
      for i := 0; i < 10; i++ {
        fmt.Printf("%02x", arr[i])
        if (i+1)%4 == 0 || i == 9 {
            fmt.Printf("\n")
        } else {
            fmt.Printf(" ")
        }
    }

      for i := 0; i < 10; i++ {
        b := arr[i]
          if b >= 32 && b <= 126 {
            fmt.Printf("%c", b)
        } else {
            fmt.Printf(".")
        }
    }
    fmt.Printf("\n")
}
