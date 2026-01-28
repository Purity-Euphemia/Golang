import "errors"

func countAbove(nums []int, limit int) int {
    count := 0
    for _, n := range nums {
        if n > limit {
            count++
        }
    }
    return count
}
