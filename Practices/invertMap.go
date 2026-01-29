func invertMap(m map[string]string) map[string]string {
    result := map[string]string{}
    for k, v := range m {
        result[v] = k
    }
    return result
}
