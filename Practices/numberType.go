func numberType(n int) string {
    if n > 0 {
        return "Positive"
    } else if n < 0 {
        return "Negative"
    }
    return "Zero"
}
