func hoursSince(t time.Time) int {
    return int(time.Since(t).Hours())
}
