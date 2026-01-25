func countDown(n int) []int {
    result := []int{}
    for i := n; i >= 0; i-- {
        result = append(result, i)
    }
    return result
}
