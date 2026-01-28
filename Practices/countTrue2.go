func countTrue(bools []bool) int {
    count := 0
    for _, b := range bools {
        if b {
            count++
        }
    }
    return count
}
