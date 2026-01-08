func countEvenOdd(nums []int) (int, int) {
    even, odd := 0, 0
    for _, n := range nums {
        if n%2 == 0 {
            even++
        } else {
            odd++
        }
    }
    return even, odd
}
