func keyExists(m map[string]int, key string) bool {
    _, ok := m[key]
    return ok
}
