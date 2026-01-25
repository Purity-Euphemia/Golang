func removeNegatives(nums []int) []int {
    result := []int{}
    for _, n := range nums {
        if n >= 0 {
            result = append(result, n)
        }
    }
    return result
}
