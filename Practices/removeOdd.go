func removeOdd(nums []int) []int {
    result := []int{}
    for _, n := range nums {
        if n%2 == 0 {
            result = append(result, n)
        }
    }
    return result
}
