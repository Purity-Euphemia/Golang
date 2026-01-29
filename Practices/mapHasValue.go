func mapHasValue(m map[string]int, value int) bool {
    for _, v := range m {
        if v == value {
            return true
        }
    }
    return false
}
