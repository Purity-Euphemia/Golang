func repeat(s string, n int) string {
    r := ""
    for i := 0; i < n; i++ { r += s }
    return r
}
