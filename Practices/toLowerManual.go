func toLowerManual(s string) string {
    result := ""
    for _, ch := range s {
        if ch >= 'A' && ch <= 'Z' {
            ch = ch + 32
        }
        result += string(ch)
    }
    return result
}
