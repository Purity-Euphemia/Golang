func minSlice(nums []int) int {
    min := nums[0]
    for _, n := range nums { if n < min { min = n } }
    return min
}
