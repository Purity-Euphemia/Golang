func allEven(nums []int) bool {
    for _, n := range nums {
        if n%2 != 0 {
            return false
        }
    }
    return true
}
