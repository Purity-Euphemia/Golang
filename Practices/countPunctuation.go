func countPunctuation(s string) int {
    count := 0
    for _, ch := range s {
        if ch == '.' || ch == ',' || ch == '!' || ch == '?' {
            count++
        }
    }
    return count
}
