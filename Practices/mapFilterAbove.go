func mapFilterAbove(m map[string]int, limit int) map[string]int {
    result := map[string]int{}
    for k, v := range m {
        if v > limit {
            result[k] = v
        }
    }
    return result
}
