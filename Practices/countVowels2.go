func countVowels(s string) int {
    c := 0
    for _, ch := range s {
        switch ch {
        case 'a','e','i','o','u','A','E','I','O','U':
            c++
        }
    }
    return c
}
