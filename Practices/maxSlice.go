func maxSlice(nums []int) int {
    max := nums[0]
    for _, n := range nums { if n > max { max = n } }
    return max
}
