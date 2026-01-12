func containsA(s string) bool {
    for _, c := range s {
        if c == 'a' || c == 'A' { return true }
    }
    return false
}
