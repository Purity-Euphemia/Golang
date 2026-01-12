func toLower(s string) string {
    result := ""
    for _, c := range s {
        if c >= 'A' && c <= 'Z' { c += 32 }
        result += string(c)
    }
    return result
}
