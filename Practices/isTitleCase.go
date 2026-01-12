func isTitleCase(s string) bool {
    if len(s) == 0 {
        return false
    }
    return s[0] >= 'A' && s[0] <= 'Z'
}
