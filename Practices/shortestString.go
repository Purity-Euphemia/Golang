func shortestString(words []string) string {
    shortest := words[0]
    for _, w := range words {
        if len(w) < len(shortest) {
            shortest = w
        }
    }
    return shortest
}
