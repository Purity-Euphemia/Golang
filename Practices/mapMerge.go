func mapMerge(a, b map[string]int) map[string]int {
    result := map[string]int{}
    for k, v := range a {
        result[k] = v
    }
    for k, v := range b {
        result[k] = v
    }
    return result
}
