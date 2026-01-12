func countDigitsInString(s string) int {
    count := 0
    for _, ch := range s {
        if ch >= '0' && ch <= '9' {
            count++
        }
    }
    return count
}
