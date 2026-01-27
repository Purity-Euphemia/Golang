func removeAtIndex(nums []int, index int) []int {
    return append(nums[:index], nums[index+1:]...)
}
