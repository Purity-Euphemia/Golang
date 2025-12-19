package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	n := 0
	for _, c := range os.Args[1] {
		if c < '0' || c > '9' {
			return
		}
		n = n*10 + int(c-'0')
	}

	if n <= 1 {
		return
	}

	first := true
	div := 2

	for n > 1 {
		if n%div == 0 {
			if !first {
				fmt.Print("*")
			}
			fmt.Print(div)
			first = false
			n /= div
		} else {
			div++
		}
	}
	fmt.Println()
}
