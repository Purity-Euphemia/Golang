func clamp(n, min, max int) int {
    if n < min {
        return min
    }
    if n > max {
        return max
    }
    return n
}
