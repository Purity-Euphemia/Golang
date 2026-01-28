func anyTrue(bools []bool) bool {
    for _, b := range bools {
        if b {
            return true
        }
    }
    return false
}
