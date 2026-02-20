package main

import "fmt"

func main() {
	var a int
	fmt.Print("Enter your number ")
	fmt.Scanln(&a)
	fmt.Println("The number ", a)
}