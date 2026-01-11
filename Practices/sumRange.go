func sumRange(start, end int) int {
    sum := 0
    for i := start; i <= end; i++ {
        sum += i
    }
    return sum
}
