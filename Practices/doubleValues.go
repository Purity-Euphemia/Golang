func doubleValues(arr []int) {
    for i := range arr {
        arr[i] *= 2
    }
}

func main() {
    nums := []int{1, 2, 3}
    doubleValues(nums)
    fmt.Println(nums) // [2 4 6]
}
