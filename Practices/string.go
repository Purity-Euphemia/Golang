package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func (p Person) Greet() string {
	return "Hello" + p.Name
}

func main() {
p := Person{Name: "John", Age: 30}
	fmt.Println(p)

}