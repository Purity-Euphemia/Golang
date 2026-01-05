func reverseString(s string) string {
    reversed := ""
    for _, char := range s {
        reversed = string(char) + reversed
    }
    return reversed
}
