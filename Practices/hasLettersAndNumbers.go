func hasLettersAndNumbers(s string) bool {
    hasLetter, hasDigit := false, false
    for _, ch := range s {
        if ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' {
            hasLetter = true
        }
        if ch >= '0' && ch <= '9' {
            hasDigit = true
        }
    }
    return hasLetter && hasDigit
}
