func countGreaterThanAverage(nums []int) int {
    sum := 0
    for _, n := range nums {
        sum += n
    }
    avg := sum / len(nums)

    count := 0
    for _, n := range nums {
        if n > avg {
            count++
        }
    }
    return count
}
