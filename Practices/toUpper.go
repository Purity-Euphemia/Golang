func toUpper(s string) string {
    result := ""
    for _, c := range s {
        if c >= 'a' && c <= 'z' { c -= 32 }
        result += string(c)
    }
    return result
}
