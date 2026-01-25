func squareSlice(nums []int) []int {
    for i := range nums {
        nums[i] *= nums[i]
    }
    return nums
}
