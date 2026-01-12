func countDivisibleBy5(nums []int) int {
    count := 0
    for _, n := range nums {
        if n%5 == 0 {
            count++
        }
    }
    return count
}
