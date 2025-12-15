package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	seen := make(map[rune]bool)

	for _, ch := range s1 {
		if seen[ch] {
			continue
		}

		for _, ch2 := range s2 {
			if ch == ch2 {
				fmt.Print(string(ch))
				seen[ch] = true
				break
			}
		}
	}

	fmt.Println()
}
