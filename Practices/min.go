func min(nums []int) int {
    smallest := nums[0]
    for _, num := range nums {
        if num < smallest {
            smallest = num
        }
    }
    return smallest
}
