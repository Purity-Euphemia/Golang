func mapValues(m map[string]int) []int {
    values := []int{}
    for _, v := range m {
        values = append(values, v)
    }
    return values
}
