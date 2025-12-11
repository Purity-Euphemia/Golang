package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	pattern := os.Args[1]
	text := os.Args[2]

	if text == "" {
		return
	}

	if len(pattern) < 3 || pattern[0] != '[' || pattern[len(pattern)-1] != ']' {
		return
	}

	content := pattern[1 : len(pattern)-1]
	if content == "" {
		return
	}

	parts := strings.Split(content, "|")

	words := strings.Fields(text)

	matches := []string{}

	for _, word := range words {
		for _, p := range parts {
			if p != "" && strings.Contains(word, p) {
				matches = append(matches, word)
				break 
			}
		}
	}

	if len(matches) == 0 {
		return
	}

	for i, m := range matches {
		fmt.Printf("%d:%s\n", i+1, m)
	}
}
