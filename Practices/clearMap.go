func clearMap(m map[string]int) {
    for k := range m {
        delete(m, k)
    }
}
