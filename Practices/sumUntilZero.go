import "errors"

func sumUntilZero(nums []int) int {
    sum := 0
    for _, n := range nums {
        if n == 0 {
            break
        }
        sum += n
    }
    return sum
}
