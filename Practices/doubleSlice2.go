func doubleSlice2(nums []int) []int {
    for i := range nums { nums[i] *= 2 }
    return nums
}
