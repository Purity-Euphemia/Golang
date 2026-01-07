func naturalNumbers(n int) []int {
    nums := []int{}
    for i := 1; i <= n; i++ {
        nums = append(nums, i)
    }
    return nums
}
