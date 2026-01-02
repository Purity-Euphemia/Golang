func increment(x *int) {
    *x = *x + 1
}

func main() {
    num := 5
    increment(&num)
    fmt.Println(num) // 6
}
