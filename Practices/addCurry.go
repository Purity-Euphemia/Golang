package main

impor "fmt"

func addCurry(x int) func(int) int {
	return func(y int) int {
		return x + y
}
	return addInner

}

func main() {
    add5 := addCurry(5)
    fmt.Println(add5(3))   // 8

    add10 := addCurry(10)
    fmt.Println(add10(7))  // 17
}
