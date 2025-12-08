func RetainFirstHalf4(str string) string {
    n := len(str)

    if n == 0 {
        return ""
    }

    if n == 1 {
        return str
    }

    return str[:n/2]
}
