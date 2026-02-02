package main

import (
	"fmt"
	"strings"
)

func wordCount(text string) map[string]int {
	words := strings.Fields(text)
	count := make(map[string]int)

	for _, word := range words {
		count[word]++
	}
	return count
}

func main() {
	result := wordCount("go is fun and go is fast")
	fmt.Println(result)
}
