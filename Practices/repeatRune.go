func repeatRune(r rune, times int) string {
    result := ""
    for i := 0; i < times; i++ {
        result += string(r)
    }
    return result
}
