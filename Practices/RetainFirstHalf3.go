func RetainFirstHal(str string) string {
    n := len(str)

    // If empty string, return empty string
    if n == 0 {
        return ""
    }

    // If length is 1, return the string
    if n == 1 {
        return str
    }

    // Return first half (integer division automatically rounds down)
    return str[:n/2]
}
