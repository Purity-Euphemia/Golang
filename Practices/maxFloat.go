func maxFloat(nums []float64) float64 {
    max := nums[0]
    for _, n := range nums {
        if n > max {
            max = n
        }
    }
    return max
}
