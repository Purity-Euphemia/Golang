func splitEvenOdd(nums []int) ([]int, []int) {
    evens := []int{}
    odds := []int{}
    for _, n := range nums {
        if n%2 == 0 {
            evens = append(evens, n)
        } else {
            odds = append(odds, n)
        }
    }
    return evens, odds
}
