package main

import "fmt"

func main() {
	age := 35
	name := "love"
	
	//print

	fmt.Println("hello")
	fmt.Println("world \n")
	fmt.Println("goodbye everyone \n")

	//println
	fmt.Println("hello everyone")
	fmt.Println("goodbye everyone")
	fmt.Println("my age is", age, "and my name is", name)

	
	// printf(formatted string)
	fmt.Printf("my age is %v and my name is %v \n", age, name)
	fmt.Printf("my age is %q and my name is %q \n", age, name)
	fmt.Printf("age is of type %!" \n, age)
	fmt.Printf("you scored %f points \n", 225.55)
	fmt.Printf("you scored %nf points \n", 225.55)

	//Sprintf(save formatted string)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age, name)
	fmt.Println("the saved string is:", str)



}