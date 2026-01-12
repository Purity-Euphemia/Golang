func containsVowel(s string) bool {
    for _, ch := range s {
        switch ch {
        case 'a','e','i','o','u','A','E','I','O','U':
            return true
        }
    }
    return false
}
