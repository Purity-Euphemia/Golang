func countOdd(nums []int) int {
    count := 0
    for _, n := range nums {
        if n%2 != 0 {
            count++
        }
    }
    return count
}
