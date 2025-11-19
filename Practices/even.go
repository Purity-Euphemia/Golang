package main

import "fmt"

func even() {
	var a int

	fmt.Println("Enter first number: ")
	fmt.Scan(&a)

	if a%2 == 0 {
		fmt.Println("even")
	}else {
		fmt.Println("odd")

}

}

func main() {
    even()

	

}