func countLetter(s string, letter rune) int {
    count := 0
    for _, ch := range s {
        if ch == letter {
            count++
        }
    }
    return count
}
