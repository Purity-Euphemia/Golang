func countOccurrences(nums []int, target int) int {
    count := 0
    for _, num := range nums {
        if num == target {
            count++
        }
    }
    return count
}
