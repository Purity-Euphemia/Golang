func averageFloat(nums []float64) float64 {
    sum := 0.0
    for _, n := range nums {
        sum += n
    }
    return sum / float64(len(nums))
}
