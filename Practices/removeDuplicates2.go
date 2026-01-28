func removeDuplicates(nums []int) []int {
    seen := map[int]bool{}
    result := []int{}
    for _, n := range nums {
        if !seen[n] {
            seen[n] = true
            result = append(result, n)
        }
    }
    return result
}
