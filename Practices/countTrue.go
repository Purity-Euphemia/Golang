func countTrue(values []bool) int {
    count := 0
    for _, v := range values {
        if v {
            count++
        }
    }
    return count
}
