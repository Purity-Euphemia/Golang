func stats(nums []int) (int, float64) {
    sum := 0
    for _, v := range nums {
        sum += v
    }
    avg := float64(sum) / float64(len(nums))
    return sum, avg
}

func main() {
    nums := []int{10, 20, 30, 40}
    sum, avg := stats(nums)
    fmt.Println("Sum:", sum)
    fmt.Println("Average:", avg)
}
