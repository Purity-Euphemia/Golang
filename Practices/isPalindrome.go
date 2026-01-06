func isPalindrome(s string) bool {
    reversed := ""
    for _, ch := range s {
        reversed = string(ch) + reversed
    }
    return s == reversed
}
