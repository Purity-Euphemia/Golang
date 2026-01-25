func mapKeys(m map[string]int) []string {
    keys := []string{}
    for k := range m {
        keys = append(keys, k)
    }
    return keys
}
