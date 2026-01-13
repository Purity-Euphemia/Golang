func countOdd(nums []int) int {
    c := 0
    for _, n := range nums { if n%2 != 0 { c++ } }
    return c
}
