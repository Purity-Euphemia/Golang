func longestString(words []string) string {
    longest := ""
    for _, w := range words {
        if len(w) > len(longest) {
            longest = w
        }
    }
    return longest
}
