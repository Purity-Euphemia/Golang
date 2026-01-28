import "errors"

func replaceNegative(nums []int) []int {
    for i, n := range nums {
        if n < 0 {
            nums[i] = 0
        }
    }
    return nums
}
