func containsSpace(s string) bool {
    for _, ch := range s {
        if ch == ' ' {
            return true
        }
    }
    return false
}
