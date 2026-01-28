func totalCharacters(words []string) int {
    count := 0
    for _, w := range words {
        count += len(w)
    }
    return count
}
