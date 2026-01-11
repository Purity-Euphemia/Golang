func isLowerCase(s string) bool {
    for _, ch := range s {
        if ch < 'a' || ch > 'z' {
            return false
        }
    }
    return true
}
