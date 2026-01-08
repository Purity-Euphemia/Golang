func countUppercase(s string) int {
    count := 0
    for _, ch := range s {
        if ch >= 'A' && ch <= 'Z' {
            count++
        }
    }
    return count
}
