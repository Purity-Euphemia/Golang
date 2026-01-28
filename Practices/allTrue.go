func allTrue(bools []bool) bool {
    for _, b := range bools {
        if !b {
            return false
        }
    }
    return true
}
