func commonElements(a, b []int) []int {
    m := make(map[int]bool)
    result := []int{}

    for _, x := range a {
        m[x] = true
    }
    for _, y := range b {
        if m[y] {
            result = append(result, y)
        }
    }
    return result
}
