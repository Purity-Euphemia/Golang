func getEvens(nums []int) []int {
    evens := []int{}
    for _, v := range nums {
        if v%2 == 0 {
            evens = append(evens, v)
        }
    }
    return evens
}
