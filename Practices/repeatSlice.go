func repeatSlice(nums []int, times int) []int {
    result := []int{}
    for i := 0; i < times; i++ {
        result = append(result, nums...)
    }
    return result
}
