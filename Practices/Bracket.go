package main

import (
	"fmt"
	"os"
)

func isBalanced(s string) bool {
	stack := []rune{}

	for _, c := range s {
		switch c {
		case '(', '[', '{':
			stack = append(stack, c)

		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if (c == ')' && top != '(') ||
				(c == ']' && top != '[') ||
				(c == '}' && top != '{') {
				return false
			}
		}
	}

	return len(stack) == 0
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		return
	}

	for _, arg := range args {
		if isBalanced(arg) {
			fmt.Println("OK")
		} else {
			fmt.Println("Error")
		}
	}
}
