func allNonEmpty(words []string) bool {
    for _, w := range words {
        if w == "" {
            return false
        }
    }
    return true
}
