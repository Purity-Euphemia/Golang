package main

import "fmt"

func main() {
	arr := [3]string{"Go", "Java", "Python"}
	for i, v := range arr {
		fmt.Println(i, v)
	}
}
