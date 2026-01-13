func multiplyAll(nums []int) int {
    p := 1
    for _, n := range nums { p *= n }
    return p
}
