func countSpaces(s string) int {
    count := 0
    for _, ch := range s {
        if ch == ' ' {
            count++
        }
    }
    return count
}
