package main

import "fmt"

func isPalindrome(s string) bool {
	runes := []rune(s)
	i, j := 0, len(runes)-1

	for i < j {
		if runes[i] != runes[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("madam")) // true
	fmt.Println(isPalindrome("hello")) // false
}
