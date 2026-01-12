func containsNegative(nums []int) bool {
    for _, n := range nums {
        if n < 0 {
            return true
        }
    }
    return false
}
