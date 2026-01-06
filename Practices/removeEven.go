func removeEven(nums []int) []int {
    result := []int{}
    for _, num := range nums {
        if num%2 != 0 {
            result = append(result, num)
        }
    }
    return result
}
