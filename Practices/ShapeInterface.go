package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width, height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func main() {
	var s Shape = Rectangle{10, 5}
	fmt.Println(s.Area())
}
