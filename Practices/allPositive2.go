func allPositive(nums []int) bool {
    for _, n := range nums {
        if n <= 0 {
            return false
        }
    }
    return true
}
