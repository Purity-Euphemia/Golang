func contains(nums []int, n int) bool {
    for _, v := range nums { if v == n { return true } }
    return false
}
