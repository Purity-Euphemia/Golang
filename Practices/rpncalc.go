package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isOperator(s string) bool {
	return s == "+" || s == "-" || s == "*" || s == "/" || s == "%"
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error")
		return
	}

	input := os.Args[1]
	tokens := strings.Fields(input)

	stack := []int{}

	for _, token := range tokens {
		if isOperator(token) {
			if len(stack) < 2 {
				fmt.Println("Error")
				return
			}

			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var res int
			switch token {
			case "+":
				res = a + b
			case "-":
				res = a - b
			case "*":
				res = a * b
			case "/":
				if b == 0 {
					fmt.Println("Error")
					return
				}
				res = a / b
			case "%":
				if b == 0 {
					fmt.Println("Error")
					return
				}
				res = a % b
			}

			stack = append(stack, res)
		} else {
			n, err := strconv.Atoi(token)
			if err != nil {
				fmt.Println("Error")
				return
			}
			stack = append(stack, n)
		}
	}

	if len(stack) != 1 {
		fmt.Println("Error")
		return
	}

	fmt.Println(stack[0])
}
