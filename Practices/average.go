func average(nums []int) float64 {
    sum := 0
    for _, num := range nums {
        sum += num
    }
    return float64(sum) / float64(len(nums))
}
